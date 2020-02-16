package gophers

import (
	"context"
	"github.com/graph-gophers/graphql-go"
	"go-graphql-comparison/models"
	"strconv"
)

type mutation struct{}

func (r *mutation) CreateTodo(ctx context.Context, args struct{ Data *createTodoInput }) (*todoResolver, error) {

	nextId := 1
	length := len(models.TodoSlice)
	if length > 0 {
		nextId = models.TodoSlice[length-1].Id + 1
	}

	newTodo := &models.Todo{
		Id:        nextId,
		Text:      args.Data.Text,
		Completed: args.Data.Completed,
	}
	models.TodoSlice = append(models.TodoSlice, newTodo)
	return &todoResolver{todo: newTodo}, nil
}

func (r *mutation) ToggleTodo(args struct{ Id graphql.ID }) (*toggleTodoOutput, error) {
	id, err := strconv.Atoi(string(args.Id))
	if err != nil {
		return nil, err
	}

	for _, todo := range models.TodoSlice {
		if todo.Id == id {
			todo.Completed = !todo.Completed
			return &toggleTodoOutput{
				ok:   true,
				todo: todo,
			}, nil
		}
	}

	return &toggleTodoOutput{ok: false}, nil
}
