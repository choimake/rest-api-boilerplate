package authenticator

import (
	"fmt"
	"net/http"
	"rest-api-boilerplate/internal/adapter/jwt"
	"rest-api-boilerplate/internal/adapter/lib"
	"rest-api-boilerplate/internal/adapter/response"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString, err := lib.BearerTokenFromHeader(r)
		if err != nil {
			message := fmt.Sprintf("unauthorized: %v", err)
			res := response.NewErrorResponse(message)
			lib.ToBody(w, http.StatusUnauthorized, res)
			return
		}

		if _, err = jwt.Parser(tokenString); err != nil {
			message := fmt.Sprintf("unauthorized: %v", err)
			res := response.NewErrorResponse(message)
			lib.ToBody(w, http.StatusUnauthorized, res)
			return
		}

		next.ServeHTTP(w, r)
	})
}
