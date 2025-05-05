package main

import (
	"cleaning-app/db"
	"cleaning-app/handlers"
	"cleaning-app/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	// Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,PATCH",
		AllowHeaders: "Content-Type, Authorization",
	}))

	db.Connect()

	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)
	app.Get("/services", handlers.GetAllServices)
	app.Get("/reviews", handlers.GetAllReviews)

	app.Post("/orders", middleware.Protected(), handlers.CreateOrder)
	app.Get("/orders", middleware.Protected(), handlers.GetMyOrders)
	app.Delete("/orders/:id", middleware.Protected(), handlers.DeleteOrder)
	app.Post("/orders/:id/reviews", middleware.Protected(), handlers.CreateReview)

	app.Get("/admin/orders", middleware.Protected(), middleware.AdminOnly(), handlers.GetAllOrders)
	app.Patch("/admin/orders/:id", middleware.Protected(), middleware.AdminOnly(), handlers.UpdateOrderStatus)
	app.Delete("/admin/orders/:id", middleware.Protected(), middleware.AdminOnly(), handlers.DeleteOrder)

	app.Listen(":8000")
}
