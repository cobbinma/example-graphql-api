package postgres

import (
	"context"
	"fmt"
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
	panic("implement me")
}

func (p *postgres) UpdateMenuItems(ctx context.Context, items []*models.MenuItem) ([]*models.MenuItem, error) {
	panic("implement me")
}
