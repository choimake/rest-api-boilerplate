package stub

import (
	"context"
	"rest-api-boilerplate/internal/entity/todo"
)

type todoRepository struct {
}

var _ todo.Repository = (*todoRepository)(nil)

var todos = []todo.Todo{
	{
		Id:       1,
		UserId:   1,
		Content:  "john first todo",
		MediaUrl: "https://example.com/1.jpg",
	},
	{
		Id:       2,
		UserId:   1,
		Content:  "john second todo",
		MediaUrl: "https://example.com/2.jpg",
	},
	{
		Id:       3,
		UserId:   2,
		Content:  "jane first todo",
		MediaUrl: "https://example.com/3.jpg",
	},
}

func NewTodoRepository() todo.Repository {
	return &todoRepository{}
}

func (r todoRepository) FetchByUserId(_ context.Context, userId int) ([]todo.Todo, error) {

	var ts []todo.Todo
	for _, t := range todos {
		if t.UserId == userId {
			ts = append(ts, t)
		}
	}

	return ts, nil
}

func (r todoRepository) CreateTodo(_ context.Context, todo *todo.Todo) error {

	id := len(todos) + 1
	todo.Id = id

	todos = append(todos, *todo)
	return nil
}
