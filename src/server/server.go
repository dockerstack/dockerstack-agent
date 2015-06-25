package main

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"github.com/shirou/gopsutil/mem"
	"logging"
)

func ServerMemory() {

	config, err := toml.LoadFile("agent.toml")
	if err != nil {
		fmt.Println(err.Error())
	} else {

		apikey := config.Get("main.apikey").(int64)

		logging.AppEvent("User API Key:", string(apikey))

		v, _ := mem.VirtualMemory()

		fmt.Println(v)
		// return v

	}
}

func ProcessList() {

}

func DiskSpace() {

}

func main() {

	ServerMemory()
}
