package deployer

import (
	"fmt"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func getCurrentPath() string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	return path
}

// ExecDeployScript
func ExecDeployScript() string {
	output, err := exec.Command(getCurrentPath() + "/deploy.sh").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(output)
}
