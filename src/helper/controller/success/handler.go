package controller_success

import "github.com/gofiber/fiber/v2"

type ReturnFn func(c *fiber.Ctx, message string, obj ...interface{}) error

func ReturnOk(c *fiber.Ctx, message string, obj ...interface{}) error {
	return c.Status(200).JSON(fiber.Map{"status": "OK", "message": message, "data": obj})
}

func ReturnCreated(c *fiber.Ctx, message string, obj ...interface{}) error {
	return c.Status(201).JSON(fiber.Map{"status": "CREATED", "message": message, "data": obj})
}

func ReturnAccepted(c *fiber.Ctx, message string, obj ...interface{}) error {
	return c.Status(202).JSON(fiber.Map{"status": "ACCEPTED", "message": message, "data": obj})
}
