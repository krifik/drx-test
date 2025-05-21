package controller

import "github.com/gofiber/fiber/v2"

type ProductController interface {
	FindAll(c *fiber.Ctx) error
	Add(c *fiber.Ctx) error
}
