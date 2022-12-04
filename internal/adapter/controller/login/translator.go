package login

import (
	"rest-api-boilerplate/internal/adapter/jwt"
	"rest-api-boilerplate/internal/usecase/login"
	"strconv"
)

func InputTranslator(r *Request) *login.Input {
	i := &login.Input{
		Id:       r.UserId,
		Password: r.Password,
	}
	return i
}

func outputTranslator(o *login.Output) (*Response, error) {

	claims := jwt.Claims{
		Subject:   strconv.Itoa(o.UserId),
		FirstName: o.FirstName,
		LastName:  o.LastName,
		Email:     o.Email,
	}

	accessToken, err := jwt.Generator(claims)
	if err != nil {
		return nil, err
	}

	res := Response{
		AccessToken: accessToken,
	}

	return &res, nil

}
