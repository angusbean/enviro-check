package routes

import (
	"github.com/angusbean/enviro-check/app/controllers"
	"github.com/angusbean/enviro-check/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for POST method:
	route.Post("/request-weather", middleware.JWTProtected(), controllers.RequestWeather) // request weather information
	route.Post("/user/sign/out", middleware.JWTProtected(), controllers.UserSignOut)      // de-authorization user
	route.Post("/token/renew", middleware.JWTProtected(), controllers.RenewTokens)        // renew Access & Refresh tokens

	// Routes for PUT method:
	//route.Put("/book", middleware.JWTProtected(), controllers.UpdateBook) // update one book by ID

	// Routes for DELETE method:
	//route.Delete("/book", middleware.JWTProtected(), controllers.DeleteBook) // delete one book by ID
}
