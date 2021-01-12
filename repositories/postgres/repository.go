package postgres

import (
	"context"
	"github.com/cobbinma/example-graphql-api/models"
)

var _ models.Repository = (*postgres)(nil)

type postgres struct{}

func NewPostgres(config *Config) (models.Repository, error) {
	return &postgres{}, nil
}

func (p *postgres) MenuItems(ctx context.Context) ([]*models.MenuItem, error) {
	panic("implement me")
}

func (p *postgres) UpdateMenuItems(ctx context.Context, items []*models.MenuItem) ([]*models.MenuItem, error) {
	panic("implement me")
}
