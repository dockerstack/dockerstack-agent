package server

import (
	"config"
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/mem"
	"logging"
)

type ServerData struct {
	Apikey string `json:"apikey"`
	Data   struct {
		Available int `json:"available"`
		Total     int `json:"total"`
		Used      int `json:"used"`
	} `json:"data"`
	Serverip string `json:"serverip"`
}

type ProcessData struct {
	Apikey string `json:"apikey"`
	Data   struct {
		Process []struct {
			Name string `json:"name"`
			Path string `json:"path"`
			Size string `json:"size"`
		} `json:"process"`
	} `json:"data"`
	Serverip string `json:"serverip"`
}

func ServerMemory() (*ServerData, error) {

	data := &ServerData{}
	v, _ := mem.VirtualMemory()

	data.Data.Total = int(v.Total / 1024)
	data.Data.Available = int(v.Available / 1024)
	data.Data.Used = int(v.Used / 1024)
	m, _ := json.Marshal(data)

	fmt.Println(string(m))

	return data, nil
}

func ProcessList() {

}

func DiskSpace() {

}
