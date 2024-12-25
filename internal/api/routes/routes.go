package routes

import (
	"hello/internal/adapters/handlers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, userHandler *handlers.UserHandler) {
	// User routes
	e.POST("/users", userHandler.CreateUser)
	e.GET("/users/:id", userHandler.GetUser)
	e.GET("/users", userHandler.GetAllUsers)
	e.PUT("/users/:id", userHandler.UpdateUser)
	e.DELETE("/users/:id", userHandler.DeleteUser)
}
