package graphqlgo

import (
	"github.com/graphql-go/graphql"
	"log"
)

var Schema *graphql.Schema

func init() {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    query,
		Mutation: mutation,
	})
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	Schema = &schema
}
