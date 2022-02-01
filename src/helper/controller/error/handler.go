package controller_error

import "github.com/gofiber/fiber/v2"

type ReturnFn func(c *fiber.Ctx, message string, err error) error

func ReturnInternalServerError(c *fiber.Ctx, message string, err error) error {
	return c.Status(500).JSON(fiber.Map{"status": "INTERNAL SERVER ERROR", "message": message, "data": fiber.Map{"error": err.Error()}})
}

func ReturnBadRequest(c *fiber.Ctx, message string, err error) error {
	return c.Status(400).JSON(fiber.Map{"status": "BAD REQUEST", "message": message, "data": fiber.Map{"error": err.Error()}})
}

func ReturnUnauthorized(c *fiber.Ctx, message string, err error) error {
	return c.Status(401).JSON(fiber.Map{"status": "UNAUTHORIZED", "message": message, "data": fiber.Map{"error": err.Error()}})
}

func ReturnForbidden(c *fiber.Ctx, message string, err error) error {
	return c.Status(403).JSON(fiber.Map{"status": "FORBIDDEN", "message": message, "data": fiber.Map{"error": err.Error()}})
}

func ReturnNotFound(c *fiber.Ctx, message string, err error) error {
	return c.Status(404).JSON(fiber.Map{"status": "NOT FOUND", "message": message, "data": fiber.Map{"error": err.Error()}})
}
