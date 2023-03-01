package main

import (
	"fmt"
	"os"
	"restAPI/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	fmt.Println(os.Getenv("SVC1_URL"))
	fmt.Println(os.Getenv("SVC2_URL"))
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)
	app.Listen(":3000")
}
