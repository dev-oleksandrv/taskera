package utils

import (
	"errors"
	"github.com/dev-oleksandrv/taskera/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"strings"
	"time"
)

var (
	signingMethod = jwt.SigningMethodHS256
)

type UserClaims struct {
	ID    uuid.UUID
	Email string
	jwt.RegisteredClaims
}

func GenerateJWTToken(id uuid.UUID, email string) string {
	token := jwt.NewWithClaims(signingMethod, UserClaims{
		ID:    id,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})
	signedToken, _ := token.SignedString([]byte(config.AppConfig.Auth.Secret))
	return signedToken
}

func ParseJWTToken(jwtToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(jwtToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.Auth.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		if claims.ExpiresAt.Before(time.Now()) {
			return nil, errors.New("token is expired")
		}
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func ExtractBearerToken(ctx *gin.Context) (string, error) {
	header := ctx.Request.Header.Get("Authorization")
	if header == "" {
		return "", errors.New("no authorization header")
	}

	parts := strings.SplitN(header, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", errors.New("invalid authorization header")
	}

	return parts[1], nil
}
