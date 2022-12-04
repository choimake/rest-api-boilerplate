package login

import (
	"fmt"
	"net/http"
	"rest-api-boilerplate/internal/adapter/lib"
	"rest-api-boilerplate/internal/adapter/response"
	"rest-api-boilerplate/internal/usecase/login"
)

var _ Controller = (*controller)(nil)

type Controller interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	usecase login.UseCase
}

func NewLoginController(useCase login.UseCase) Controller {
	return &controller{
		usecase: useCase,
	}
}

func (l controller) Handle(w http.ResponseWriter, r *http.Request) {

	req, err := Encode(r)
	if err != nil {
		res := response.NewErrorResponse(err.Error())
		lib.ToBody(w, http.StatusBadRequest, res)
		return
	}

	if err := req.Validate(); err != nil {
		res := response.NewErrorResponse(err.Error())
		lib.ToBody(w, http.StatusBadRequest, res)
		return
	}

	i := InputTranslator(req)

	ctx := r.Context()
	o, err := l.usecase.Handle(ctx, i)
	if err != nil {
		message := fmt.Sprintf("fail to login: %v", err)
		res := response.NewErrorResponse(message)
		lib.ToBody(w, http.StatusInternalServerError, res)
		return
	}

	res, err := outputTranslator(o)
	if err != nil {
		res := response.NewErrorResponse(err.Error())
		lib.ToBody(w, http.StatusInternalServerError, res)
		return
	}

	lib.ToBody(w, http.StatusOK, res)
}
