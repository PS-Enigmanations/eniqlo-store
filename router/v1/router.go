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
	Product     *ProductRouter
	Customer    *CustomerRouter
	Staff       *StaffRouter
	Transaction *TransactionRouter
}

func NewV1Router(ctx context.Context, pool *pgxpool.Pool) *v1Router {
	return &v1Router{
		Product:     NewProductRouter(ctx, pool),
		Customer:    NewCustomerRouter(ctx, pool),
		Staff:       NewStaffRouter(ctx, pool),
		Transaction: NewTransactionRouter(ctx, pool),
	}
}

func (v *v1Router) Load(router *gin.Engine, m middleware.Middleware) {
	v1 := router.Group("/v1")
	{
		// Customer api endpoint
		customer := v1.Group("/customer").Use(m.Auth.MustAuthenticated())
		{
			customer.GET("/", v.Customer.Controller.SearchCustomer)
			customer.POST("/register", v.Customer.Controller.Register)
		}

		//Staff api endpoint
		staff := v1.Group("/staff")
		{
			staff.POST("/register", v.Staff.Controller.Register)
			staff.POST("/login", v.Staff.Controller.Login)
		}

		// Product api endpoint
		product := v1.Group("/product")
		{
			product.GET("/customer", v.Product.Controller.SearchProducts)

			product.Use(m.Auth.MustAuthenticated())
			product.GET("/", v.Product.Controller.Index)
			product.POST("/", v.Product.Controller.CreateProduct)
			product.PUT("/:id", v.Product.Controller.UpdateProduct)
			product.DELETE("/:id", v.Product.Controller.DeleteProduct)

			checkout := product.Group("/checkout")
			{
				checkout.POST("/", v.Transaction.Controller.Checkout)
				checkout.GET("/history", v.Transaction.Controller.SearchTransaction)
			}
		}
	}
}
