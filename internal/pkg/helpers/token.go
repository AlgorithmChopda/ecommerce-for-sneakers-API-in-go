package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(userId, role int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":   userId,
			"role": role,
			// in env time
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	tokenString, err := token.SignedString([]byte(jwtSecretKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
