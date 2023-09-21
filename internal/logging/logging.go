package logging

import (
	"log"
	"os"
)

func initLogging() {
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
}
