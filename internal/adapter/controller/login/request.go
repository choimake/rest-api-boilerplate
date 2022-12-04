package login

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"net/http"
	"rest-api-boilerplate/internal/adapter/lib"
)

type Request struct {
	Password string `json:"password"`
	UserId   int    `json:"userId"`
}

func (req Request) Validate() error {
	err := validation.ValidateStruct(&req,
		validation.Field(&req.UserId, validation.Required),
		validation.Field(&req.Password, validation.Required, validation.Length(1, 128)))
	return err
}

func Encode(r *http.Request) (*Request, error) {

	req := &Request{}
	err := lib.BodyToEntity(r, req)
	if err != nil {
		return nil, err
	}

	return req, nil
}
