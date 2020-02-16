package main

import (
	gqlgenhandler "github.com/99designs/gqlgen/handler"
	"github.com/graph-gophers/graphql-go/relay"
	graphqlgohandler "github.com/graphql-go/handler"
	"go-graphql-comparison/gophers"
	"go-graphql-comparison/gqlgen"
	"go-graphql-comparison/graphqlgo"
	"net/http"
)

func main() {
	handleGraphqlGo()
	handleGqlgen()
	handleGophers()
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

func handleGophers() {
	gophersHandler := &relay.Handler{Schema: gophers.Schema}
	// There's no playground, so we use gqlgen's playground.
	// You can copy 'github.com/99designs/gqlgen@v0.10.2/handler/playground.go' and use it.
	playgroundHandler := gqlgenhandler.Playground("GraphQL playground", "/gophers")

	http.HandleFunc("/gophers", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			playgroundHandler.ServeHTTP(w, r)
		case "POST":
			gophersHandler.ServeHTTP(w, r)
		}
	})
}
