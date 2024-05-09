package router_v1

import (
	"context"
	"enigmanations/eniqlo-store/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type V1Router interface {
	Load(r *gin.Engine, m middleware.Middleware)
}

type v1Router struct {
	Product *ProductRouter
}

func NewV1Router(ctx context.Context, pool *pgxpool.Pool) *v1Router {
	return &v1Router{
		Product: NewProductRouter(ctx, pool),
	}
}

func (v *v1Router) Load(router *gin.Engine, m middleware.Middleware) {
	v1 := router.Group("/v1")
	{
		product := v1.Group("/product")
		{
			product.GET("/customer", v.Product.Controller.SearchProducts)
		}
	}
}
