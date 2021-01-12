package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/cobbinma/example-graphql-api/graph/generated"
	"github.com/cobbinma/example-graphql-api/models"
)

func (r *mutationResolver) UpdateMenuItems(ctx context.Context, items []*models.MenuItemInput) ([]*models.MenuItem, error) {
	menuItems := []*models.MenuItem{}
	for i := range items {
		item := items[i].MenuItem(r.timeClient.now())
		menuItems = append(menuItems, &item)
	}

	return r.repository.UpdateMenuItems(ctx, menuItems)
}

func (r *queryResolver) MenuItems(ctx context.Context) ([]*models.MenuItem, error) {
	return r.repository.MenuItems(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
