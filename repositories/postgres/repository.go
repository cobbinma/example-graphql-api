package postgres

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/cobbinma/example-graphql-api/models"
	"github.com/golang-migrate/migrate/v4"
	pgmigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var _ models.Repository = (*postgres)(nil)

type postgres struct {
	db  *sqlx.DB
	log *zap.SugaredLogger
}

func NewPostgres(config *Config) (models.Repository, error) {
	log := config.log
	db, err := sqlx.Connect("postgres", config.pgURL.String())
	if err != nil {
		return nil, fmt.Errorf("could not connect to postgres : %w", err)
	}

	driver, err := pgmigrate.WithInstance(db.DB, &pgmigrate.Config{})
	if err != nil {
		return nil, fmt.Errorf("could not create database driver : %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		return nil, fmt.Errorf("error instantiating migrate : %w", err)
	}

	version, dirty, _ := m.Version()
	log.Infof("Database version %d, dirty %t", version, dirty)

	log.Infof("Starting migration")
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return nil, fmt.Errorf("an error occurred while syncing the database.. %w", err)
	}

	return &postgres{db: db}, nil
}

func (p *postgres) MenuItems(ctx context.Context) ([]*models.MenuItem, error) {
	sql, args, err := sq.Select("*").
		From("menu_items").
		Where("available_at > NOW()").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("could not construct sql : %w", err)
	}

	items := []*models.MenuItem{}
	err = p.db.Select(&items, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("could not get menu items from postgres : %w", err)
	}

	return items, nil
}

func (p *postgres) UpdateMenuItems(ctx context.Context, items []*models.MenuItem) ([]*models.MenuItem, error) {
	tx, err := p.db.Beginx()
	if err != nil {
		return nil, fmt.Errorf("could not begin transaction : %w", err)
	}

	if err := insertAndUpdateItems(ctx, tx, items); err != nil {
		return nil, fmt.Errorf("could not insert and update items : %w", err)
	}

	var itemIDs []string
	for i := range items {
		itemIDs = append(itemIDs, items[i].ID)
	}

	if err := removeExcludedItems(ctx, tx, itemIDs); err != nil {
		return nil, fmt.Errorf("could not remove excluded items : %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return items, nil
}

func insertAndUpdateItems(ctx context.Context, tx *sqlx.Tx, items []*models.MenuItem) error {
	if len(items) < 1 {
		return nil
	}

	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Insert("menu_items").
		Columns("id", "status", "available_at")

	for i := range items {
		builder = builder.Values(items[i].ID, items[i].Status, items[i].AvailableAt)
	}
	builder = builder.Suffix("ON CONFLICT (id) DO UPDATE SET status = EXCLUDED.status, available_at = EXCLUDED.available_at")

	sql, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("could not build sql statement : %w", err)
	}
	_, err = tx.Exec(sql, args...)
	if err != nil {
		return fmt.Errorf("could not execute : %w", err)
	}
	return nil
}

func removeExcludedItems(ctx context.Context, tx *sqlx.Tx, itemIDs []string) error {
	sql, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Delete("menu_items").
		Where(sq.And{
			sq.NotEq{"id": itemIDs},
		}).ToSql()
	if err != nil {
		return fmt.Errorf("could not build sql statement : %w", err)
	}
	_, err = tx.Exec(sql, args...)
	if err != nil {
		return fmt.Errorf("could not execute : %w", err)
	}
	return nil
}
