package server

import (
	"github.com/VAISHAKH-GK/msg-grid-auth/internal/router"
	"github.com/gofiber/fiber/v2"
)

// Running server
func Run(port string) {
	// Fiber router
	var app = fiber.New()

	// Handling '/auth' routes
	router.AuthRouter(app.Group("/auth"))

	// starting server and listening to specified port
	app.Listen(":" + port)
}
