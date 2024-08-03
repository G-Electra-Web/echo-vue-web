package api

import (
	"git.nothr.in/nothr/g-electra/pkg/handlers"
	"github.com/labstack/echo"
)

// Initialise Routes
func InitialiseRoutes(e *echo.Echo) {

	// Group API
	apiRoute := e.Group("/api")

	// Handle SignUp
	apiRoute.POST("/signup", handlers.SignUp)
}
