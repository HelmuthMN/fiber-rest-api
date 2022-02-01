package controller_helper

import "github.com/gofiber/fiber/v2"

type GetObjectFromPostRequestFn func(c *fiber.Ctx, obj interface{}) error

func GetObjectFromPostRequest(c *fiber.Ctx, obj interface{}) error {
	return c.BodyParser(obj)
}
