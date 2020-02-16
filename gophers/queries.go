package gophers

import (
	"context"
	"go-graphql-comparison/models"
)

type query struct{}

func (r *query) Todos(ctx context.Context) ([]*todoResolver, error) {
	todoResolvers := make([]*todoResolver, len(models.TodoSlice))
	for i, todo := range models.TodoSlice {
		todoResolvers[i] = &todoResolver{todo: todo}
	}
	return todoResolvers, nil
}
