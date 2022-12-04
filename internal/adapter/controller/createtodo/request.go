package createtodo

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"mime/multipart"
	"net/http"
	"rest-api-boilerplate/internal/adapter/jwt"
	"rest-api-boilerplate/internal/adapter/lib"
	"strconv"
)

type Request struct {
	UserId      int
	Content     string
	Media       multipart.File
	MediaHeader *multipart.FileHeader
}

func (req Request) Validate() error {
	err := validation.ValidateStruct(&req,
		validation.Field(&req.UserId, validation.Required),
		validation.Field(&req.Content, validation.Required, validation.Length(1, 256)),
		validation.Field(&req.Media),
		validation.Field(&req.MediaHeader))
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

	content := r.FormValue("content")

	m, h, err := r.FormFile("media")
	if err != http.ErrMissingFile {
		return nil, err
	}

	req := &Request{}

	req.UserId = userId
	req.Content = content
	req.Media = m
	req.MediaHeader = h

	return req, nil
}
