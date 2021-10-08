package deployer

import (
	"github.com/gofiber/fiber"
	log "github.com/sirupsen/logrus"
)

func ReceiveRequestPost(c *fiber.Ctx) error {
	log.Info(string(c.Body()))
	return c.Status(200).Send([]byte(ExecDeployScript()))
}
func ReceiveRequestGet(c *fiber.Ctx) error {

	return c.Status(200).Send([]byte(LastDeployTime()))
}
