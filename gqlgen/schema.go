package gqlgen

import "go-graphql-comparison/gqlgen/gen"

var Schema = gen.NewExecutableSchema(gen.Config{Resolvers: &Resolver{}})

type Resolver struct{}

func (r *Resolver) Mutation() gen.MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() gen.QueryResolver {
	return &queryResolver{r}
}
