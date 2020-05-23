package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/AlexKoppara/Hermes/graph/generated"
	"github.com/AlexKoppara/Hermes/graph/model"
)

func (r *entityResolver) FindMenuItemByID(ctx context.Context, id string) (*model.MenuItem, error) {
	return &model.MenuItem{
		ID: id,
	}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
