package main

import (
	"context"
	"enigmanations/eniqlo-store/config"
	"enigmanations/eniqlo-store/middleware"
	"enigmanations/eniqlo-store/pkg/database"
	"enigmanations/eniqlo-store/pkg/env"
	routes "enigmanations/eniqlo-store/router"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

	// Prepare middleware
	middleware := middleware.NewMiddleware(pool)

	// Prepare router
	router := gin.New()

	// Register routes
	routes.RegisterRouter(ctx, pool, router, middleware)

	// Run the server
	appServeAddr := ":" + fmt.Sprint(cfg.AppPort)
	fmt.Printf("Serving on http://localhost:%s\n", fmt.Sprint(cfg.AppPort))
	log.Fatalf("%v", http.ListenAndServe(appServeAddr, router))
}
