package v1

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type V1Router interface {
	Load(r *gin.Engine)
}

type v1Router struct {
	Customer     *CustomerRouter
}

func NewV1Router(ctx context.Context, pool *pgxpool.Pool) *v1Router {
    return &v1Router{
        Customer:     NewCustomerRouter(ctx, pool),
    }
}

func (v *v1Router) Load(router *gin.Engine) {
	// @see https://gin-gonic.com/docs/examples/grouping-routes/
	v1 := router.Group("/v1")
	{
		// Customer api endpoint
		customer := v1.Group("/customer")
		{
			customer.POST("/register", v.Customer.Controller.CustomerRegister)
		}
	}
}