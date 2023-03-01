package routes

import (
	"restAPI/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("user", controller.CreateUser)

	app.Post("post", controller.CreatePost)
}
