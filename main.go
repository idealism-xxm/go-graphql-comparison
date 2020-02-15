package main

import (
	"github.com/graphql-go/handler"
	"go-graphql-comparison/graphqlgo"
	"net/http"
)

func main() {
	graphqlGoHandler := handler.New(&handler.Config{
		Schema:     graphqlgo.Schema,
		Pretty:     true,
		Playground: true,
	})

	http.Handle("/graphqlGo", graphqlGoHandler)
	http.ListenAndServe(":8080", nil)
}
