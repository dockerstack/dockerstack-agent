package logging

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/pelletier/go-toml"
	"os"
)

func init() {

	config, err := toml.LoadFile("agent.toml")

	if err != nil {
		fmt.Println("Error: ", err.Error())
	} else {

		logfile := config.Get("main.log").(string)

		logFile, err := os.OpenFile(logfile, os.O_SYNC|os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

		if err != nil {
			panic(err)
		}

		log.SetOutput(logFile)
		log.SetLevel(log.DebugLevel)
		log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	}
}

func Critical(message string, value string) {
	log.Warnf(message, value)
}

func ApiKey(message string, key string) {
	log.Infof(message, key)
}
func ResourceAlert(message string, alert string) {
	log.Info(message, alert)
}

func CodeProfilingInfo(message string, data string) {
	log.Debug(message, data)
}

func AppEvent(message string, event string) {
	log.Info(message, event)
}
