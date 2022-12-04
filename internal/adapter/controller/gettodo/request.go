package gettodo

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"net/http"
	"rest-api-boilerplate/internal/adapter/jwt"
	"rest-api-boilerplate/internal/adapter/lib"
	"strconv"
)

type Request struct {
	UserId int
}

func (req Request) Validate() error {
	err := validation.ValidateStruct(&req,
		validation.Field(&req.UserId, validation.Required))
	return err
}

func Encode(r *http.Request) (*Request, error) {

	tokenString, err := lib.BearerTokenFromHeader(r)
	if err != nil {
		return nil, err
	}

	claims, err := jwt.Parser(tokenString)
	if err != nil {
		return nil, err
	}

	userId, err := strconv.Atoi(claims.Subject)

	if err != nil {
		return nil, err
	}

	req := &Request{
		UserId: userId,
	}

	return req, nil
}
