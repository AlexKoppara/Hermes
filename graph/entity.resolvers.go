package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/AlexKoppara/Hermes/graph/generated"
	"github.com/AlexKoppara/Hermes/graph/model"
)

func (r *entityResolver) FindMenuItemByID(ctx context.Context, id string) (*model.MenuItem, error) {
	var menuItem *model.MenuItem
	switch id {
	case "menu_item_1":
		menuItem = &model.MenuItem{
			ID: "menu_item_1",
		}
	case "menu_item_2":
		menuItem = &model.MenuItem{
			ID: "menu_item_2",
		}
	case "menu_item_3":
		menuItem = &model.MenuItem{
			ID: "menu_item_3",
		}
	}
	return menuItem, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
