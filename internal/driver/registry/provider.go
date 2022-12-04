package registry

import (
	"rest-api-boilerplate/internal/adapter/controller/createtodo"
	"rest-api-boilerplate/internal/adapter/controller/gettodo"
	"rest-api-boilerplate/internal/adapter/controller/login"
	"rest-api-boilerplate/internal/driver/stub"
	createtodo2 "rest-api-boilerplate/internal/usecase/createtodo"
	gettodo2 "rest-api-boilerplate/internal/usecase/gettodo"
	login2 "rest-api-boilerplate/internal/usecase/login"
)

var provider = &Provider{}

type Provider struct {
	LoginController      login.Controller
	GetTodoController    gettodo.Controller
	CreateTodoController createtodo.Controller
}

func InitProvider(r *Resource) {
	provider.LoginController = initLoginController(r)
	provider.GetTodoController = initGetTodoController(r)
	provider.CreateTodoController = initCreateTodoController(r)
}

func initLoginController(_ *Resource) login.Controller {
	r := stub.NewUserRepository()
	u := login2.NewUseCase(r)
	return login.NewLoginController(u)
}

func initGetTodoController(_ *Resource) gettodo.Controller {
	r := stub.NewTodoRepository()
	u := gettodo2.NewUseCase(r)
	return gettodo.NewGetTodoController(u)
}

func initCreateTodoController(_ *Resource) createtodo.Controller {
	r := stub.NewTodoRepository()
	u := createtodo2.NewUseCase(r)
	return createtodo.NewCreateTodoController(u)
}

// GetProvider returns the Provider.
func GetProvider() *Provider {
	return provider
}
