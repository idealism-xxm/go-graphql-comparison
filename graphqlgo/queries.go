package graphqlgo

import (
	"github.com/graphql-go/graphql"
	"go-graphql-comparison/models"
)

var query = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"todos": todos(),
	},
})

func todos() *graphql.Field {
	return &graphql.Field{
		Name: "todos",
		Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(todo))),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return models.TodoSlice, nil
		},
		Description: "list all todos",
	}
}
