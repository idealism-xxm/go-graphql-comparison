package gen

import (
	"context"
	"go-graphql-comparison/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Todo() TodoResolver {
	return &todoResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, data CreateTodoInput) (*models.Todo, error) {
	panic("not implemented")
}
func (r *mutationResolver) ToggleTodo(ctx context.Context, id string) (*ToggleTodoOutput, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	panic("not implemented")
}

type todoResolver struct{ *Resolver }

func (r *todoResolver) ID(ctx context.Context, obj *models.Todo) (string, error) {
	panic("not implemented")
}
