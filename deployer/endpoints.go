package deployer

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func ReceiveRequestPost(c *fiber.Ctx) error {
	log.Info(string(c.Body()))

	if !IsValidMessage(c) {
		return c.Status(403).Send([]byte("Forbiden"))
	}

	return c.Status(200).Send([]byte(ExecDeployScript()))
}
func ReceiveRequestGet(c *fiber.Ctx) error {

	return c.Status(200).Send([]byte(LastDeployTime()))
}
