package deployer

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	log "github.com/sirupsen/logrus"
)

var lastdeploy = time.Now()

func getCurrentPath() string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	return path
}

// ExecDeployScript
func ExecDeployScript() string {
	output, err := exec.Command("/bin/sh", getCurrentPath()+"/deploy.sh").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	lastdeploy = time.Now()
	return string(output)
}

func LastDeployTime() string {
	return fmt.Sprint("Last deploy time: ", lastdeploy)
}
