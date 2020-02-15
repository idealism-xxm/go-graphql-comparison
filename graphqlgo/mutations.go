package graphqlgo

import (
	"github.com/graphql-go/graphql"
	"go-graphql-comparison/models"
	"strconv"
)

var mutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createTodo": createTodo(),
		"toggleTodo": toggleTodo(),
	},
})

func createTodo() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewNonNull(todo),
		Args: graphql.FieldConfigArgument{
			"data": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(createTodoInput),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			data := p.Args["data"].(map[string]interface{})
			text := data["text"].(string)
			completed := data["completed"].(bool)

			nextId := 1
			length := len(models.TodoSlice)
			if length > 0 {
				nextId = models.TodoSlice[length-1].Id + 1
			}

			newTodo := &models.Todo{
				Id:        nextId,
				Text:      text,
				Completed: completed,
			}
			models.TodoSlice = append(models.TodoSlice, newTodo)

			return newTodo, nil
		},
		Description: "create a new todo item and return it",
	}
}

func toggleTodo() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewNonNull(toggleTodoOutput),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, err := strconv.Atoi(p.Args["id"].(string))
			if err != nil {
				return nil, err
			}

			for _, todo := range models.TodoSlice {
				if todo.Id == id {
					todo.Completed = !todo.Completed
					return &toggleTodoOutputStruct{
						Ok:   true,
						Todo: todo,
					}, nil
				}
			}

			return &toggleTodoOutputStruct{Ok: false}, nil
		},
		Description: "toggle todo item and return it if exists",
	}
}
