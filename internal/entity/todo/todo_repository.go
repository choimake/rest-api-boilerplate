package todo

import "context"

type Repository interface {
	FetchByUserId(ctx context.Context, userId int) ([]Todo, error)
	CreateTodo(ctx context.Context, todo *Todo) error
}
