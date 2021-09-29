package config

import (
	"fmt"
	"log"
	"strconv"
	
	"github.com/BurntSushi/toml"
)

var (
	conf Config
)

type Config struct {
	HTTP  map[string]int
}

// Load configurations from config.toml file
func Load() {
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		panic(err)
	}
}

//GetHTTPPort returns http port
func GetHTTPPort() string {
	return fmt.Sprintf(":%d", conf.HTTP["port"])
}
