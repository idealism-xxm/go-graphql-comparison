package gqlgen

import (
	"context"
	"go-graphql-comparison/gqlgen/gen"
	"go-graphql-comparison/models"
	"strconv"
)

func (r *Resolver) Todo() gen.TodoResolver {
	return &todoResolver{r}
}

type todoResolver struct{ *Resolver }

func (r *todoResolver) ID(ctx context.Context, obj *models.Todo) (string, error) {
	return strconv.Itoa(obj.Id + 1), nil
}
