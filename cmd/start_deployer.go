package cmd

import (
	"github.com/gofiber/fiber"
	"github.com/spf13/cobra"
	"github.com/tmota900/ss-deployer/deployer"
	"github.com/tmota900/ss-deployer/config"
	"fmt"
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
var StartDeployer = &cobra.Command{
	Use:   "deployer",
	Short: "start listning to / route",
	Run: func(cmd *cobra.Command, args []string) {
		// load configs
		config.Load()

		fmt.Println("yeeeeee")
		// configure endpoints
		router := setHandlers()
		// start server
		router.Listen(config.GetHTTPPort())
	},
}
