package config

import (
	"fmt"
	
	"github.com/BurntSushi/toml"
)

var (
	conf Config
)

type Config struct {
	HTTP  map[string]string
}

// Load configurations from config.toml file
func Load() {
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		//panic(err)
	}
}

//GetHTTPPort returns http port
func GetHTTPPort() string {
	if(conf.HTTP["port"] == ""){
		return ":1337"
	}
	return fmt.Sprintf(":%d", conf.HTTP["port"])
}
