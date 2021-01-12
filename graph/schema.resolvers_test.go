package graph_test

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/bradleyjkemp/cupaloy"
	"github.com/cobbinma/example-graphql-api/graph"
	"github.com/cobbinma/example-graphql-api/graph/generated"
	"github.com/cobbinma/example-graphql-api/models"
	"github.com/cobbinma/example-graphql-api/repositories/fakerepository"
	"testing"
	"time"
)

var date = time.Date(1992, 5, 1, 0, 0, 0, 0, time.UTC)

func Test_UpdateMenuItems(t *testing.T) {
	repository := fakerepository.NewFake()
	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.NewResolver(repository, graph.WithFixedTime(date))})))

	var resp struct {
		UpdateMenuItems []struct {
			ID          string
			Status      string
			AvailableAt string
		}
	}
	c.MustPost(`mutation { updateMenuItems(items: [{id: "fd361dae-97ee-4847-9a3d-1bcbc506c2dd", status: HIDDEN}]) { id, status, availableAt } }`, &resp)

	cupaloy.SnapshotT(t, resp)
}

func Test_UpdateMenuItems_Null(t *testing.T) {
	repository := fakerepository.NewFake()
	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.NewResolver(repository)})))

	var resp struct {
		UpdateMenuItems []struct {
			ID          string
			Status      string
			AvailableAt string
		}
	}
	c.MustPost(`mutation { updateMenuItems(items: []) { id, status, availableAt } }`, &resp)

	cupaloy.SnapshotT(t, resp)
}

func Test_MenuItems(t *testing.T) {
	repository := fakerepository.NewFake()
	_, err := repository.UpdateMenuItems(context.Background(), []*models.MenuItem{
		{
			ID:     "def78745-d30c-4876-86a9-8067b76a2275",
			Status: models.ItemStatusHidden,
		},
		{
			ID:          "3c556919-753f-42dc-bbfa-90eb31e22792",
			Status:      models.ItemStatusUnavailable,
			AvailableAt: &date,
		},
	})
	if err != nil {
		panic(fmt.Sprintf("did not expect an error, got '%s'", err))
	}

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.NewResolver(repository, graph.WithFixedTime(date))})))

	var resp struct {
		MenuItems []struct {
			ID          string
			Status      string
			AvailableAt string
		}
	}
	c.MustPost(`{ menuItems { id, status, availableAt } }`, &resp)

	cupaloy.SnapshotT(t, resp)
}

func Test_MenuItems_Null(t *testing.T) {
	repository := fakerepository.NewFake()
	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.NewResolver(repository)})))

	var resp struct {
		MenuItems []struct {
			ID          string
			Status      string
			AvailableAt string
		}
	}
	c.MustPost(`{ menuItems { id, status, availableAt } }`, &resp)

	cupaloy.SnapshotT(t, resp)
}
