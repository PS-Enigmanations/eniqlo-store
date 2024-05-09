package router

import (
	"context"
	"enigmanations/eniqlo-store/middleware"
	v1 "enigmanations/eniqlo-store/router/v1"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterRouter(ctx context.Context, pool *pgxpool.Pool, router *gin.Engine, m middleware.Middleware) {
	v1Route := v1.NewV1Router(ctx, pool)
	v1Route.Load(router, m)
}
