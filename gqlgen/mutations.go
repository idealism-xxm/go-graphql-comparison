package gqlgen

import (
	"context"
	"go-graphql-comparison/gqlgen/gen"
	"go-graphql-comparison/models"
	"strconv"
)

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, data gen.CreateTodoInput) (*models.Todo, error) {
	nextId := 1
	length := len(models.TodoSlice)
	if length > 0 {
		nextId = models.TodoSlice[length-1].Id + 1
	}

	completed := false
	if data.Completed != nil {
		completed = *data.Completed
	}

	newTodo := &models.Todo{
		Id:        nextId,
		Text:      data.Text,
		Completed: completed,
	}
	models.TodoSlice = append(models.TodoSlice, newTodo)

	return newTodo, nil
}
func (r *mutationResolver) ToggleTodo(ctx context.Context, id string) (*gen.ToggleTodoOutput, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	for _, todo := range models.TodoSlice {
		if todo.Id == idInt {
			todo.Completed = !todo.Completed
			return &gen.ToggleTodoOutput{
				Ok:   true,
				Todo: todo,
			}, nil
		}
	}

	return &gen.ToggleTodoOutput{Ok: false}, nil
}
