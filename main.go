package main

import (
	"hello/internal/adapters/handlers"
	"hello/internal/adapters/repositories/memory"
	"hello/internal/api/routes"
	"hello/internal/core/services"
	"log"

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

	// Initialize dependencies
	userRepo := memory.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Setup routes
	routes.SetupRoutes(e, userHandler)

	// Start server
	if err := e.Start(":1323"); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
