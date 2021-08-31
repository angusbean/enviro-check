package routes

import (
	"github.com/angusbean/enviro-check/app/controllers"
	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for GET method:
	//route.Get("/example", controllers.GetBooks)   // get list of all books
	//route.Get("/book/:id", controllers.GetBook) // get one book by ID
	route.Get("/test", controllers.Test)

	// Routes for POST method:
	route.Post("/user/sign/up", controllers.UserSignUp) // register a new user
	route.Post("/user/sign/in", controllers.UserSignIn) // auth, return Access & Refresh tokens
}
