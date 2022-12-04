package chi

import (
	"github.com/go-chi/chi/v5"
	"rest-api-boilerplate/internal/adapter/authenticator"
	"rest-api-boilerplate/internal/driver/middleware"
	"rest-api-boilerplate/internal/driver/registry"
)

func NewRouter(p *registry.Provider) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.CorsHeaderHandler)

	// without authentication
	r.Group(func(r chi.Router) {
		r.Post("/login", p.LoginController.Handle)
	})

	// authentication required
	r.Group(func(r chi.Router) {
		r.Use(authenticator.Authenticate)
		r.Get("/todos", p.GetTodoController.Handle)
		r.Post("/todos", p.CreateTodoController.Handle)
	})

	return r
}
