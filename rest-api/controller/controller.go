package controller

import (
	"context"
	"fmt"
	"restAPI/services"
	svcspb "restAPI/svcspb"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(ctx *fiber.Ctx) error {
	fmt.Println("Create user rest")
	body := new(svcspb.User)
	err := ctx.BodyParser(body)
	if err != nil {
		return err
	}
	res, err := services.UserSrv().CreateUser(context.Background(), &svcspb.CreateUserRequest{
		Username: body.Username,
	})
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(res.User)
}

func CreatePost(ctx *fiber.Ctx) error {
	fmt.Println("Create post rest")
	body := new(svcspb.Post)
	err := ctx.BodyParser(body)
	if err != nil {
		return err
	}
	res, err := services.PostSrv().CreatePost(context.Background(), &svcspb.CreatePostRequest{
		Post: &svcspb.Post{
			Title:   body.Title,
			Content: body.Content,
		},
	})
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(res.Post)
}
