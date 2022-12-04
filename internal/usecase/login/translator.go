package login

import "rest-api-boilerplate/internal/entity/user"

func outputTranslator(user *user.User) *Output {
	return &Output{
		UserId:    user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}
