package main

import (
	gqlgenhandler "github.com/99designs/gqlgen/handler"
	graphqlgohandler "github.com/graphql-go/handler"
	"go-graphql-comparison/gqlgen"
	"go-graphql-comparison/graphqlgo"
	"net/http"
)

func main() {
	handleGraphqlGo()
	handleGqlgen()
	http.ListenAndServe(":8080", nil)
}

func handleGraphqlGo() {
	graphqlGoHandler := graphqlgohandler.New(&graphqlgohandler.Config{
		Schema:     graphqlgo.Schema,
		Pretty:     true,
		Playground: true,
	})
	http.Handle("/graphqlGo", graphqlGoHandler)
}

func handleGqlgen() {
	gqlgenHandler := gqlgenhandler.GraphQL(gqlgen.Schema)
	playgroundHandler := gqlgenhandler.Playground("GraphQL playground", "/gqlgen")

	http.HandleFunc("/gqlgen", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			playgroundHandler.ServeHTTP(w, r)
		case "POST":
			gqlgenHandler.ServeHTTP(w, r)
		}
	})
}
