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
		customer := v1.Group("/customer")
		{
			customer.GET("/", m.Auth.MustAuthenticated(), v.Customer.Controller.SearchCustomer)
			customer.POST("/register", m.Auth.MustAuthenticated(), v.Customer.Controller.Register)
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
			product.GET("/", m.Auth.MustAuthenticated(), v.Product.Controller.Index)
			product.POST("/", m.Auth.MustAuthenticated(), v.Product.Controller.CreateProduct)
			product.PUT("/:id", m.Auth.MustAuthenticated(), v.Product.Controller.UpdateProduct)
			product.DELETE("/:id", m.Auth.MustAuthenticated(), v.Product.Controller.DeleteProduct)
			product.GET("/customer", v.Product.Controller.SearchProducts)

			checkout := product.Group("/checkout")
			{
				checkout.GET("/history", m.Auth.MustAuthenticated(), v.Transaction.Controller.SearchTransaction)
				checkout.POST("/", m.Auth.MustAuthenticated(), v.Transaction.Controller.Checkout)
			}
		}
	}
}
