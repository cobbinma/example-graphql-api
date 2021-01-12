package repositorytest

import (
	"context"
	"fmt"
	"github.com/bradleyjkemp/cupaloy"
	"github.com/cobbinma/example-graphql-api/models"
	"testing"
	"time"
)

var date = time.Date(3000, 5, 1, 0, 0, 0, 0, time.UTC)

type TestCase struct {
	Name string
	Test func(t *testing.T)
}

func TestSuite(repository models.Repository) []TestCase {
	snap := cupaloy.New(cupaloy.SnapshotSubdirectory("../repositorytest/.snapshots"))
	return []TestCase{
		{
			Name: "UpdateMenuItems_Add_Unavailable_And_Hidden_Items",
			Test: func(t *testing.T) {
				ctx := context.Background()
				defer cleanUp(ctx, repository)

				items, err := repository.UpdateMenuItems(ctx, []*models.MenuItem{
					{ID: "fd361dae-97ee-4847-9a3d-1bcbc506c2dd", Status: models.ItemStatusHidden},
					{ID: "30d087ef-2945-40d4-ba28-6bd697d8fb4e", Status: models.ItemStatusUnavailable, AvailableAt: &date},
				})
				if err != nil {
					t.Errorf("did not expect error, got '%s'", err)
					return
				}

				snap.SnapshotT(t, items)
			},
		},
		{
			Name: "UpdateMenuItems_Null",
			Test: func(t *testing.T) {
				ctx := context.Background()
				defer cleanUp(ctx, repository)

				items, err := repository.UpdateMenuItems(ctx, []*models.MenuItem{})
				if err != nil {
					t.Errorf("did not expect error, got '%s'", err)
					return
				}

				if len(items) != 0 {
					t.Errorf("expected an empty array")
				}
			},
		},
		{
			Name: "UpdateMenuItems_Overwrite_Menu_Items",
			Test: func(t *testing.T) {
				ctx := context.Background()
				defer cleanUp(ctx, repository)

				_, err := repository.UpdateMenuItems(ctx, []*models.MenuItem{
					{ID: "fd361dae-97ee-4847-9a3d-1bcbc506c2dd", Status: models.ItemStatusHidden},
					{ID: "30d087ef-2945-40d4-ba28-6bd697d8fb4e", Status: models.ItemStatusUnavailable, AvailableAt: &date},
				})
				if err != nil {
					t.Errorf("did not expect error, got '%s'", err)
					return
				}

				items, err := repository.UpdateMenuItems(ctx, []*models.MenuItem{})
				if err != nil {
					t.Errorf("did not expect error, got '%s'", err)
					return
				}

				if len(items) != 0 {
					t.Errorf("expected an empty array")
				}
			},
		},
		{
			Name: "MenuItems_Unavailable_And_Hidden_Items",
			Test: func(t *testing.T) {
				ctx := context.Background()
				defer cleanUp(ctx, repository)

				_, err := repository.UpdateMenuItems(ctx, []*models.MenuItem{})
				if err != nil {
					t.Errorf("did not expect error, got '%s'", err)
					return
				}

				items, err := repository.MenuItems(ctx)
				if err != nil {
					t.Errorf("did not expect error, got '%s'", err)
					return
				}

				if len(items) != 0 {
					t.Errorf("expected an empty array")
				}
			},
		},
		{
			Name: "MenuItems_Only_Return_Unavailable_Items_Available_In_Future",
			Test: func(t *testing.T) {
				ctx := context.Background()
				defer cleanUp(ctx, repository)

				var past = time.Date(1992, 5, 1, 0, 0, 0, 0, time.UTC)

				_, err := repository.UpdateMenuItems(ctx, []*models.MenuItem{
					{ID: "30d087ef-2945-40d4-ba28-6bd697d8fb4e", Status: models.ItemStatusUnavailable, AvailableAt: &date},
					{ID: "ea61e7a7-b302-4b00-bf46-3fc78de595cd", Status: models.ItemStatusUnavailable, AvailableAt: &past},
				})
				if err != nil {
					t.Errorf("did not expect error, got '%s'", err)
					return
				}

				items, err := repository.MenuItems(ctx)
				if err != nil {
					t.Errorf("did not expect error, got '%s'", err)
					return
				}

				snap.SnapshotT(t, items)
			},
		},
	}
}

func cleanUp(ctx context.Context, repository models.Repository) {
	_, err := repository.UpdateMenuItems(ctx, []*models.MenuItem{})
	if err != nil {
		panic(fmt.Sprintf("could not clean up repository : %s", err))
	}
}
