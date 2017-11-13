package config

import (
	"os"

	"github.com/jinzhu/configor"
)

// Config
var Config = struct {
	Host string `default:"0.0.0.0" env:"HOST"`
	Port uint   `default:"7000" env:"PORT"`
}{}

// Set environment variable config path -> export APP_CONFIG=/etc/qor/production.yml
func init() {
	if fileConfig := os.Getenv("APP_CONFIG"); len(fileConfig) > 0 {
		if err := configor.Load(&Config, fileConfig); err != nil {
			panic(err)
		}
	} else {
		if err := configor.Load(&Config); err != nil {
			panic(err)
		}
	}
}
