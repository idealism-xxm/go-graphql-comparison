package gophers

import (
	"github.com/graph-gophers/graphql-go"
	"go-graphql-comparison/models"
	"strconv"
)

type todoResolver struct {
	todo *models.Todo
}

func (r todoResolver) Id() (graphql.ID, error) {
	return graphql.ID(strconv.Itoa(r.todo.Id)), nil
}

func (r todoResolver) Text() (string, error) {
	return r.todo.Text, nil
}

func (r todoResolver) Completed() (bool, error) {
	return r.todo.Completed, nil
}
