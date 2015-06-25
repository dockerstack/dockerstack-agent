package config

import (
	_ "fmt"
	"github.com/pelletier/go-toml"
	"logging"
)

func ApiKey() (apikey string) {

	config, _ := toml.LoadFile("agent.toml")

	akey := config.Get("main.apikey").(string)
	apikey = akey

	logging.ApiKey("This is the User API Key:", apikey)
	return apikey
}

func Main(value string) (data string) {

	config, _ := toml.LoadFile("agent.toml")

	if value == "apikey" {
		data := config.Get("main.apikey").(string)

		return data
	} else if value == "apiserver" {

		data := config.Get("main.apiserver").(string)
		return data
	} else if value == "log" {
		data := config.Get("main.log").(string)
		return data
	} else if value == "interface" {
		data := config.Get("main.interface").(string)
		return data
	}

	return data
}
