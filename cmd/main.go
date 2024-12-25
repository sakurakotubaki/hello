package main

import (
	"hello/internal/adapters/handlers"
	"hello/internal/adapters/repositories/memory"
	"hello/internal/api/routes"
	"hello/internal/core/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initialize Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Initialize repositories
	userRepo := memory.NewUserRepository()

	// Initialize services
	userService := services.NewUserService(userRepo)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)

	// Setup routes
	routes.SetupRoutes(e, userHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
