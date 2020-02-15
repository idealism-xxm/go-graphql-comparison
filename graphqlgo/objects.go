package graphqlgo

import (
	"errors"
	"github.com/graphql-go/graphql"
	"go-graphql-comparison/models"
)

var todo = graphql.NewObject(graphql.ObjectConfig{
	Name: "Todo",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Name: "id",
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
				if todo, ok := p.Source.(*models.Todo); ok {
					return todo.Id, nil
				}
				return nil, errors.New("not supported type, expected type: `*models.Todo`")
			},
			Description: "todo id",
		},
		"text": &graphql.Field{
			Name: "text",
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
				if todo, ok := p.Source.(*models.Todo); ok {
					return todo.Text, nil
				}
				return nil, errors.New("not supported type, expected type: `*models.Todo`")
			},
			Description: "todo body text",
		},
		"completed": &graphql.Field{
			Name: "completed",
			Type: graphql.NewNonNull(graphql.Boolean),
			Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
				if todo, ok := p.Source.(*models.Todo); ok {
					return todo.Completed, nil
				}
				return nil, errors.New("not supported type, expected type: `*models.Todo`")
			},
			Description: "is it completed?",
		},
	},
	Description: "todo item",
})
