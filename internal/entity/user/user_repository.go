package user

import "context"

type Repository interface {
	FetchByIdAndPassword(ctx context.Context, id int, password string) (*User, error)
}
