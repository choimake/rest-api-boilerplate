package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

var signingKey = "signingKey"

func Generator(claims Claims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 署名
	accessToken, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func Parser(tokenString string) (*Claims, error) {

	// jwtの検証
	// Reference: https://pkg.go.dev/github.com/golang-jwt/jwt/v4#example-Parse-Hmac
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token is invalid")
}
