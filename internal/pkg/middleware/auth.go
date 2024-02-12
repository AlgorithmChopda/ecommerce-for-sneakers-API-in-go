package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/dto"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/helpers"
	"github.com/golang-jwt/jwt/v5"
)

func CheckAuth(handlerFunc func(w http.ResponseWriter, r *http.Request), access int) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			ErrorResponse(w, http.StatusUnauthorized, errors.New("authorization header not found"))
			return
		}

		tokenString = tokenString[len("Bearer "):]
		token, err := helpers.VerifyToken(tokenString)
		if err != nil {
			fmt.Println(err)
			ErrorResponse(w, http.StatusUnauthorized, errors.New("invalid token"))
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			id := int(claims["id"].(float64))
			tokenRole := int(claims["role"].(float64))

			if tokenRole != access {
				ErrorResponse(w, http.StatusUnauthorized, errors.New("invalid token"))
				return
			}

			r_copy := r.WithContext(context.WithValue(r.Context(), "token", dto.JwtToken{
				Id:   id,
				Role: tokenRole,
			}))

			handlerFunc(w, r_copy)
			return
		}

		ErrorResponse(w, http.StatusUnauthorized, errors.New("invalid token"))
	}
}
