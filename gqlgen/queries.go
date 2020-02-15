package gqlgen

import (
	"context"
	"go-graphql-comparison/models"
)

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	return models.TodoSlice, nil
}
