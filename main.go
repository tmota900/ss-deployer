package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tmota900/ss-deployer/cmd"
)

var RootCmd = &cobra.Command{}

func main() {
	RootCmd.AddCommand(cmd.StartDeployer())

	if err := RootCmd.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(-1)
	}
}
