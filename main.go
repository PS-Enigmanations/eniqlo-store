package main

import (
	"context"
	"enigmanations/eniqlo-store/config"
	"enigmanations/eniqlo-store/pkg/database"
	"enigmanations/eniqlo-store/pkg/env"
	"fmt"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	routes "enigmanations/eniqlo-store/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load .env %v\n", err)
		os.Exit(1)
	}

	cfg := config.GetConfig()

	// Shared ctx
	ctx := context.Background()

	// Connect to the database
	pgUrl := `postgres://%s:%s@%s:%d/%s?%s`
	pgUrl = fmt.Sprintf(pgUrl,
		cfg.DBUsername,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.DBParams,
	)

	pool, err := database.NewPGXPool(ctx, pgUrl, &database.PGXStdLogger{
		Logger: slog.Default(),
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	// Check reachability
	if _, err = pool.Exec(ctx, `SELECT 1`); err != nil {
		errMsg := fmt.Errorf("pool.Exec() error: %v", err)
		fmt.Println(errMsg) // or handle the error message in some other way
	}

	// Disable debug mode in production
	if env.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}

	// Prepare router
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Register routes
	routes.RegisterRouter(ctx, pool, router)

	// router.Run(fmt.Sprintf("%s:%d", cfg.AppHost, cfg.AppPort))
}
