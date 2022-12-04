package gettodo

import (
	"context"
	"rest-api-boilerplate/internal/entity/todo"
)

type UseCase interface {
	Handle(ctx context.Context, input *Input) (*Output, error)
}

type usecase struct {
	todoRepository todo.Repository
}

func NewUseCase(todoRepository todo.Repository) UseCase {
	return &usecase{
		todoRepository: todoRepository,
	}
}

func (u usecase) Handle(ctx context.Context, input *Input) (*Output, error) {

	todos, err := u.todoRepository.FetchByUserId(ctx, input.UserId)
	if err != nil {
		return nil, err
	}

	o := Output{
		Todos: todos,
	}

	return &o, nil
}
