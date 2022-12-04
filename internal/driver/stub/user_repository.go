package stub

import (
	"context"
	"errors"
	"rest-api-boilerplate/internal/entity/user"
)

type userRepository struct {
	users []user.User
}

var _ user.Repository = (*userRepository)(nil)

func NewUserRepository() user.Repository {

	users := []user.User{
		{
			Id:        1,
			Password:  "password",
			Email:     "john.due@example.com",
			FirstName: "John",
			LastName:  "Due",
		},
		{
			Id:        2,
			Password:  "password",
			Email:     "jane.due@example.com",
			FirstName: "Jane",
			LastName:  "Due",
		},
	}

	return &userRepository{
		users: users,
	}
}

func (r userRepository) FetchByIdAndPassword(_ context.Context, id int, password string) (*user.User, error) {

	for _, u := range r.users {
		if u.Id == id && u.Password == password {
			return &u, nil
		}
	}

	return nil, errors.New("user not found")
}
