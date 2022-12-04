package createtodo

import (
	"context"
	"fmt"
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

	var mediaUrl string
	if input.MediaHeader != nil {
		mediaUrl = fmt.Sprintf("https://example.com/%s", input.MediaHeader.Filename)
	}

	t := todo.Todo{
		UserId:   input.UserId,
		Content:  input.Content,
		MediaUrl: mediaUrl,
	}

	if err := u.todoRepository.CreateTodo(ctx, &t); err != nil {
		return nil, err
	}

	o := Output{
		Todo: t,
	}

	return &o, nil
}
