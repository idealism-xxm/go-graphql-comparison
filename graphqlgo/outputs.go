package graphqlgo

import (
	"errors"
	"github.com/graphql-go/graphql"
	"go-graphql-comparison/models"
)

var toggleTodoOutput = graphql.NewObject(graphql.ObjectConfig{
	Name: "ToggleTodoOutput",
	Fields: graphql.Fields{
		"ok": &graphql.Field{
			Name: "ok",
			Type: graphql.NewNonNull(graphql.Boolean),
			Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
				if todo, ok := p.Source.(*toggleTodoOutputStruct); ok {
					return todo.Ok, nil
				}
				return nil, errors.New("not supported type, expected type: `*toggleTodoOutputStruct`")
			},
			Description: "is mutation success?",
		},
		"todo": &graphql.Field{
			Name: "todo",
			Type: todo,
			Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
				if todo, ok := p.Source.(*toggleTodoOutputStruct); ok {
					return todo.Todo, nil
				}
				return nil, errors.New("not supported type, expected type: `*toggleTodoOutputStruct`")
			},
			Description: "toggled todo item if exists",
		},
	},
	Description: "todo item",
})

type toggleTodoOutputStruct struct {
	Ok   bool
	Todo *models.Todo
}
