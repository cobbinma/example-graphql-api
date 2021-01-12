package models

import (
	"context"
)

//go:generate mockgen -package=mock_models -destination=./mock/repository.go -source=repository.go
type Repository interface {
	MenuItems(ctx context.Context) ([]*MenuItem, error)
	UpdateMenuItems(ctx context.Context, items []*MenuItem) ([]*MenuItem, error)
}
