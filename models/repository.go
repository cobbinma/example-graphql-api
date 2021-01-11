package models

import (
	"context"
)

type Repository interface {
	MenuItems(ctx context.Context) ([]*MenuItem, error)
	UpdateMenuItems(ctx context.Context, items []*MenuItemInput) ([]*MenuItem, error)
}
