package routes

import (
	"github.com/angusbean/enviro-check/app/controllers"
	"github.com/angusbean/enviro-check/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group by API version
	route := a.Group("/api/v1")

	// Routes for POST method:
	route.Post("/request-weather", controllers.RequestWeather)                       // request weather information
	route.Post("/user/sign/out", middleware.JWTProtected(), controllers.UserSignOut) // de-authorization user
	route.Post("/token/renew", middleware.JWTProtected(), controllers.RenewTokens)   // renew Access & Refresh tokens
}
