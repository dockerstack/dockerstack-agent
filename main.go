package main

import (
	"bytes"
	"fmt"
	_ "github.com/gorilla/mux"
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"logging"
	"net/http"
	"server"
)

func main() {

	config, err := toml.LoadFile("agent.toml")
	if err != nil {
		fmt.Println("Error: ", err.Error())
	} else {
		apikey := config.Get("main.apikey").(int64)
		logfile := config.Get("main.log").(string)

		//r := mux.NewRouter()
		//r.HandleFunc("/", HelloWorld).Methods("GET")

		//fmt.Printf("Starting the Server on port %s", port)
		//http.ListenAndServe(":8080", r)

		response, _ := http.Get("http://www.google.com")

		body, _ := ioutil.ReadAll(response.Body)

		raw := bytes.NewBuffer(body)

		s := raw.String()
		fmt.Println(s)

		k, _ := server.ServerMemory()

		fmt.Println(k.Total)

		logging.AppEvent("Here is the LogFile Location :", logfile)

		fmt.Println(apikey, logfile)
	}
}
