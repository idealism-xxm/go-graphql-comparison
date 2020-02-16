package gophers

import "github.com/graph-gophers/graphql-go"

var Schema = graphql.MustParseSchema(schemaString, &root{})

type root struct {
	*query
	*mutation
}

// should read from schema.graphql in production
var schemaString = `
schema {
    query: Query
    mutation: Mutation
}

type Query {
    # list all todos
    todos: [Todo!]!
}

type Mutation {
    # create a new todo item and return it
    createTodo(data: CreateTodoInput!): Todo!
    # toggle todo item and return it if exists
    toggleTodo(id: ID!): ToggleTodoOutput!
}

# todo item
type Todo {
    # todo id
    id: ID!
    # todo body text
    text: String!
    # is it completed?
    completed: Boolean!
}

# toggle todo output
type ToggleTodoOutput {
    # is mutation success?
    ok: Boolean!
    # toggled todo item if exists
    todo: Todo
}

# create todo input
input CreateTodoInput {
    # todo body text
    text: String!
    # is it completed?
    completed: Boolean = false
}
`
