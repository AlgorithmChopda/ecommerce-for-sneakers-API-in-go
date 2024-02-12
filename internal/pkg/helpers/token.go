package helpers

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/apperrors"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/dto"
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

func VerifyToken(tokenString string) (*jwt.Token, error) {
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil || !token.Valid {
		fmt.Println(err)
		return nil, apperrors.UnauthorizedAccess{Message: "invalid token"}
	}

	return token, nil
}

func GetTokenData(ctx context.Context) (dto.JwtToken, error) {
	if ctx == nil {
		return dto.JwtToken{}, apperrors.UnauthorizedAccess{Message: "token data not found"}
	}

	tokenData, ok := ctx.Value("token").(dto.JwtToken)
	if !ok {
		return dto.JwtToken{}, apperrors.UnauthorizedAccess{Message: "token data not found"}
	}

	return tokenData, nil
}
