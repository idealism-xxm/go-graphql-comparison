package gophers

import (
	"go-graphql-comparison/models"
)

type toggleTodoOutput struct {
	ok   bool
	todo *models.Todo
}

func (r *toggleTodoOutput) Ok() (bool, error) {
	return r.ok, nil
}

func (r *toggleTodoOutput) Todo() (*todoResolver, error) {
	if r.todo == nil {
		return nil, nil
	}
	return &todoResolver{todo: r.todo}, nil
}
