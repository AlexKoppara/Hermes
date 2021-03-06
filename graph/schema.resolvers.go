package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/AlexKoppara/Hermes/graph/generated"
	"github.com/AlexKoppara/Hermes/graph/model"
)

func (r *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	var menuItems []*model.MenuItem
	menuItems = append(menuItems, &model.MenuItem{
		ID: "menu_item_1",
	})
	menuItems = append(menuItems, &model.MenuItem{
		ID: "menu_item_2",
	})

	var orders []*model.Order
	orders = append(orders, &model.Order{
		ID:        "order_1",
		MenuItems: menuItems,
	})
	return orders, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
