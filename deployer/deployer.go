package deployer

import (
	"fmt"
	"mime"
	"os"
	"os/exec"
	"time"
	"bytes"

	"github.com/gofiber/fiber"
	"github.com/google/go-github/github"
	log "github.com/sirupsen/logrus"
)

var (
	lastdeploy = time.Now()
	secretKey  = []byte("0123456789abcdef")
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

func IsValidMessage(c *fiber.Ctx) bool {

	_, err := ValidatePayload(c, secretKey)
	if err != nil {
		log.Error(err)
		return false
	}

	return true
}

func ValidatePayload(c *fiber.Ctx, secretToken []byte) (payload []byte, err error) {
	signature := string(c.Request().Header.Peek(github.SHA256SignatureHeader))
	if signature == "" {
		signature = string( c.Request().Header.Peek(github.SHA1SignatureHeader) )
	}

	contentType, _, err := mime.ParseMediaType( string(c.Request().Header.Peek("Content-Type")) )
	if err != nil {
		return nil, err
	}

	return github.ValidatePayloadFromBody(contentType, bytes.NewBuffer(c.Body()), signature, secretToken)
}
