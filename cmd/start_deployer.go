package cmd

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"github.com/tmota900/ss-deployer/deployer"
)

// setHandlers define http routes
func setHandlers() *fiber.App {
	// define basic gin-router
	router := fiber.New()

	router.Get("/", deployer.ReceiveRequestGet)

	router.Post("/", deployer.ReceiveRequestPost)

	return router
}

// StartDeployer starts deployer
func StartDeployer() *cobra.Command {
	var port, secret string
	c := &cobra.Command{
		Use:   "deployer",
		Short: "start listning to / route",
		Run: func(cmd *cobra.Command, args []string) {

			deployer.SetSecret(secret)

			// configure endpoints
			router := setHandlers()
			// start server
			err := router.Listen(fmt.Sprintf(":%s", port))

			if err != nil {
				log.Fatal(err)
			}
		},
	}
	c.Flags().StringVarP(&port, "port", "p", "1337", "Target port listner")
	c.Flags().StringVarP(&secret, "secret", "s", "", "Configured secret")
	return c
}
