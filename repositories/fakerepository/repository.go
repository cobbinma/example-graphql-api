package fakerepository

import (
	"context"
	"github.com/cobbinma/example-graphql-api/models"
	"time"
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
	items := []*models.MenuItem{}
	for i := range f.items {
		if f.items[i].AvailableAt == nil || f.items[i].AvailableAt.After(time.Now()) {
			items = append(items, f.items[i])
		}
	}
	return items, nil
}

func (f *fake) UpdateMenuItems(ctx context.Context, items []*models.MenuItem) ([]*models.MenuItem, error) {
	f.items = items
	return f.items, nil
}
