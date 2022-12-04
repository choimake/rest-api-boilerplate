package gettodo

import (
	"rest-api-boilerplate/internal/usecase/gettodo"
)

func InputTranslator(r *Request) *gettodo.Input {

	i := &gettodo.Input{
		UserId: r.UserId,
	}
	return i
}

func outputTranslator(o *gettodo.Output) []Todo {

	var todos []Todo

	for _, todo := range o.Todos {
		t := Todo{
			Id:       todo.Id,
			UserId:   todo.UserId,
			Content:  todo.Content,
			MediaUrl: todo.MediaUrl,
		}
		todos = append(todos, t)
	}

	return todos
}
