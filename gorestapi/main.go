package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ocg.com/go/01/lab/gorestapi/routes"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
	})

	app.Static("/public", "./public", fiber.Static{ //http://localhost:3000/public OR http://localhost:3000/public/dog.jpeg
		Compress:  true,
		ByteRange: true,
		Browse:    true,
		MaxAge:    3600,
	})

	productRouter := app.Group("/api/product")
	routes.ConfigProductRouter(&productRouter) //http://localhost:3000/api/product

	cartRouter := app.Group("/api/cart")
	routes.ConfigCartRouter(&cartRouter) //http://localhost:3000/api/cart

	categoryRouter := app.Group("/api/category")
	routes.ConfigCategoryRouter(&categoryRouter) //http://localhost:3000/api/category

	priceRouter := app.Group("/api/price")
	routes.ConfigPriceRouter(&priceRouter) //http://localhost:3000/api/price

	userRouter := app.Group("/api/user")
	routes.ConfigUserRouter(&userRouter) //http://localhost:3000/api/user

	reviewRouter := app.Group("/api/review")
	routes.ConfigReviewRouter(&reviewRouter) //http://localhost:3000/api/review
	app.Listen(":3000")
}
