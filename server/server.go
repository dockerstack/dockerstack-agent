package server

import (
	_ "bufio"
	"dockerstack/dockerstack-agent/config"
	_ "dockerstack/dockerstack-agent/logging"
	"fmt"
	ps "github.com/mitchellh/go-ps"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	_ "github.com/shirou/gopsutil/process"
	_ "io/ioutil"
	"net"
	_ "os"
	"os/exec"
	_ "reflect"
	_ "regexp"
	"strconv"
	"strings"
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

/*
Sample output
{
  "apikey": 123456,
  "serverip" : "172.27.9.56",
  "data":[{
    "name":"docker",
    "pid": 925
  },
  {
    "name": "apache2",
    "pid":1024
  }]
}
*/

type ProcessData struct {
	Apikey   string `json:"apikey"`
	Data     []Process
	Serverip string `json:"serverip"`
}

type Process struct {
	Name string `json:"name"`
	Pid  int    `json:"pid"`
}

func ServerMemory() (*ServerData, error) {

	data := &ServerData{}
	v, _ := mem.VirtualMemory()

	ief, _ := net.InterfaceByName(config.Main("interface"))

	addrs, err := ief.Addrs()
	if err != nil {
		err.Error()
	}

	data.Data.Total = int(v.Total / 1024)
	data.Data.Available = int(v.Available / 1024)
	data.Data.Used = int(v.Used / 1024)
	data.Serverip = addrs[0].String()
	data.Apikey = config.ApiKey()

	fmt.Println(ProcessList())
	return data, nil
}

//Lists Top 5 ProcessPids with Names running on the Host
func ProcessList() (*ProcessData, error) {

	k := getPids()

	fmt.Println(k)

	pids, _ := ps.Processes()

	data := new(ProcessData)

	ief, _ := net.InterfaceByName(config.Main("interface"))

	addrs, err := ief.Addrs()
	if err != nil {
		err.Error()
	}

	data.Apikey = config.ApiKey()
	data.Serverip = addrs[0].String()

	for m := 0; m < len(k); m++ {

		for _, pid := range pids {

			if k[m] == strconv.Itoa(pid.Pid()) {

				data.Data = append(data.Data, Process{Name: pid.Executable(), Pid: pid.Pid()})

			}
		}

	}

	return data, nil

}

func DiskSpace() (*disk.DiskUsageStat, error) {

	data, _ := disk.DiskUsage("/")

	return data, nil

}

//Internal Functions

//Gets Top 5 Process Id's
func getPids() []string {

	//pids := &ProcesPids{}
	cmd := "ps aux | awk '{print $2, $4, $11}' | sort -k2rn | head -n 5|awk '{print $1}'"

	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
	}

	k := strings.Split(string(out), "\n")

	fmt.Println(k[5])
	return k
}
