package middleware

import (
	"context"
	v1 "enigmanations/eniqlo-store/router/v1"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Router struct {
	ctx  context.Context
	pool *pgxpool.Pool
}

func RegisterRouter(ctx context.Context, pool *pgxpool.Pool, router *gin.Engine) {
	v1Route := v1.NewV1Router(ctx, pool)
	v1Route.Load(router)
}