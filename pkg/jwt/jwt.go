package jwt

import (
	"enigmanations/eniqlo-store/internal/staff"
	"enigmanations/eniqlo-store/pkg/env"
	"time"

	"github.com/golang-jwt/jwt"
)

const AccessTokenDuration = 8 * time.Hour // 8 hours

func GenerateAccessToken(staffID string, credential *staff.Staff) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": staffID,
		"sub": credential.Name,
		"exp": time.Now().Add(AccessTokenDuration).Unix(),
	})
	tokenString, err := token.SignedString([]byte(env.GetSecretKey()))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
