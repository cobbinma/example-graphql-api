package postgres

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/cobbinma/example-graphql-api/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var _ models.Repository = (*postgres)(nil)

type postgres struct {
	db *sqlx.DB
}

func NewPostgres(config *Config) (models.Repository, error) {
	db, err := sqlx.Connect("postgres", config.pgURL.String())
	if err != nil {
		return nil, fmt.Errorf("could not connect to postgres : %w", err)
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
	err = p.db.Select(items, sql, args)
	if err != nil {
		return nil, fmt.Errorf("could not get menu items from postgres : %w", err)
	}

	return items, nil
}

func (p *postgres) UpdateMenuItems(ctx context.Context, items []*models.MenuItem) ([]*models.MenuItem, error) {
	panic("implement me")
}
