package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const (
	AuthorizationPayloadKey = "authorization_payload"
)

func GetCurrentUser(c *gin.Context) *jwt.Token {
	return c.MustGet(AuthorizationPayloadKey).(*jwt.Token)
}
