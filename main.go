package main

import (
	"bytes"
	"config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	_ "logging"
	"net/http"
	"server"
)

func main() {

	k, _ := server.ServerMemory()

	test := config.Main("interface")

	fmt.Println(test)

	mar, _ := json.Marshal(k)

	buf := bytes.NewBuffer(mar)
	enc := json.NewEncoder(buf)
	err := enc.Encode(mar)

	if err != nil {
		panic(err)
	}

	// response, _ := http.Get("http://www.google.com")

	send, _ := http.Post("https://demo7965648.mockable.io/serverdata", "application/json", buf)

	body, _ := ioutil.ReadAll(send.Body)

	raw := bytes.NewBuffer(body)

	s := raw.String()

	fmt.Println(s)
	fmt.Println(enc)

}
