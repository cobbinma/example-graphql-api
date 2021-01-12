package fakerepository

import (
	"context"
	"github.com/cobbinma/example-graphql-api/models"
)

type fake struct {
	items []*models.MenuItem
}

var _ models.Repository = (*fake)(nil)

func NewFake() models.Repository {
	return &fake{
		items: []*models.MenuItem{},
	}
}

func (f *fake) MenuItems(ctx context.Context) ([]*models.MenuItem, error) {
	return f.items, nil
}

func (f *fake) UpdateMenuItems(ctx context.Context, items []*models.MenuItem) ([]*models.MenuItem, error) {
	f.items = items
	return f.items, nil
}
