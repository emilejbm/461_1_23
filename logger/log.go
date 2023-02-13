package logger

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	path, ok := os.LookupEnv("LOG_FILE")
	if !ok {
		fmt.Println("Location for log file not found")
		os.Exit(1)
	}

	file, err := os.OpenFile(path+"log.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Failed to create log file")
		os.Exit(1)
	}

	log.SetOutput(file)
	logLvl, ok := os.LookupEnv(("LOG_LEVEL"))

	if !ok {
		fmt.Println("Log level not found")
		os.Exit(1)
	}

	if logLvl == "1" {
		log.SetLevel(log.InfoLevel)
	} else if logLvl == "2" {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.PanicLevel) // silent mode (0)
	}
}

func DebugMsg(msg ...string) {
	log.Debug(msg)
}

func InfoMsg(msg ...string) {
	log.Info(msg)
}
