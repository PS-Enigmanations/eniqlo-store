package middleware

import (
	"context"
	"enigmanations/eniqlo-store/internal/common/auth"
	"enigmanations/eniqlo-store/internal/staff/repository"
	"enigmanations/eniqlo-store/pkg/jwt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthMiddleware struct {
	ctx  context.Context
	pool *pgxpool.Pool
}

func NewAuthMiddleware(ctx context.Context, pool *pgxpool.Pool) AuthMiddleware {
	return AuthMiddleware{ctx: ctx, pool: pool}
}

func (am AuthMiddleware) MustAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the user is authenticated
		token, err := jwt.GetTokenFromAuthHeader(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		decodedToken, err := jwt.ValidateToken(token)
		log.Println(decodedToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		staffRepo := repository.NewStaffRepository(am.pool)
		_, err = staffRepo.FindById(am.ctx, decodedToken.Uid)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		c.Header("Authorization", "Bearer "+token)
		c.Set(auth.AuthorizationPayloadKey, decodedToken)
		c.Next()
	}
}
