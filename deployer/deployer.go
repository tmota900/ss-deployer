package deployer

import (
	"bytes"
	"fmt"
	"mime"
	"os"
	"os/exec"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/go-github/v39/github"
	log "github.com/sirupsen/logrus"
	"github.com/tmota900/ss-deployer/utils"
)

var (
	lastdeploy = time.Now()
	secretKey  = []byte("")
)

func SetSecret(secret string) {
	secretKey = []byte(secret)
}

func getCurrentPath() string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	return path
}

// ExecDeployScript
func ExecDeployScript() string {
	output, err := exec.Command("/bin/bash", utils.Getenv("DEPLOY_SCRIPT_DIR", getCurrentPath()+"/deploy.sh")).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	lastdeploy = time.Now()

	log.Info(string(output))
	return string(output)
}

func LastDeployTime() string {
	return fmt.Sprint("Last deploy time: ", lastdeploy)
}

func IsValidMessage(c *fiber.Ctx) bool {

	if bytes.Compare(secretKey, []byte("")) == 0 {
		return true
	}

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
		signature = string(c.Request().Header.Peek(github.SHA1SignatureHeader))
	}

	contentType, _, err := mime.ParseMediaType(string(c.Request().Header.Peek("Content-Type")))
	if err != nil {
		return nil, err
	}

	return github.ValidatePayloadFromBody(contentType, bytes.NewBuffer(c.Body()), signature, secretToken)
}
