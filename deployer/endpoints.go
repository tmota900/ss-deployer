package deployer

import (
	"github.com/gofiber/fiber"
)


func ReceiveRequestPost(c *fiber.Ctx) error { return c.Status(200).Send([]byte("ok")) }
func ReceiveRequestGet(c *fiber.Ctx) error { return c.Status(200).Send([]byte("ok")) }