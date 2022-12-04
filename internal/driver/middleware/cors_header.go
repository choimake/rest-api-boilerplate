package middleware

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
)

func CorsHeaderHandler(next http.Handler) http.Handler {
	return middleware.RouteHeaders().Route(
		// CORS対応
		"Origin", "*", cors.Handler(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		})).Handler(next)
}
