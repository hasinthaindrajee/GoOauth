package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

func ReadConfig() Config {
	var conf Config
	if _, err := toml.DecodeFile("./config.toml", &conf); err != nil {
		fmt.Println(err)
	}
	return conf
}

type Config struct {
	Idp             Idp
	ServiceProvider ServiceProvider
}

type Idp struct {
	Host              string
	TokenEndpoint     string
	AuthorizeEndpoint string
	ClientId          string
	ClientSecret      string
}

type ServiceProvider struct {
	Callback string
	Scopes   []string
}
