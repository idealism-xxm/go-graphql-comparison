package graphqlgo

import "github.com/graphql-go/graphql"

var createTodoInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "CreateTodoInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"text": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "todo body text",
		},
		"completed": &graphql.InputObjectFieldConfig{
			Type:         graphql.Boolean,
			DefaultValue: false,
			Description:  "is it completed?",
		},
	},
	Description: "create todo input",
})
