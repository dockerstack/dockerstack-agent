package server

import (
	_ "dockerstack/dockerstack-agent/config"
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/mem"
	_ "github.com/shirou/gopsutil/process"
	_ "dockerstack/dockerstack-agent/logging"
	_ "os"
	"os/exec"
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

type ProcesPids struct {
	Process []struct {
		Pid int `json:"pid"`
	} `json:"process"`
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

func ProcessList() ([]string, error) {

	pids,_ := getPids()

	fmt.Print(pids)

	// var proc []string

	// for _, pid := range pids {

		// fmt.Println(string(pid))
		// 	p, _ := process.NewProcess(pid)

		// 	// data := &ProcessData{}

		// 	fmt.Println(p.Name())
		// 	fmt.Println(p.Exe())
		// 	fmt.Println(p.MemoryInfo())

		// 	// proc[0] = p.Name()
		// 	// proc[1] = p.Exe()
		// 	// proc[2] = p.MemoryInfo()

	// }

	return nil, nil

}

func DiskSpace() {

}

//Internal Functions

//Gets Top 5 Process Id's
func getPids() (*ProcesPids, error) {

	pids := &ProcesPids{}
	cmd := "ps aux | awk '{print $2, $4, $11}' | sort -k2rn | head -n 5|awk '{print $1}'"

	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
	}

// 	for _, dat := range pids.Process {
//
// 		for _, proc := range out {
//
// 			dat.Pid = string(proc)
// 		}
// 	}
//
	fmt.Print(string(out))

	return out

}
