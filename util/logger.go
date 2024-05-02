package util

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Logger struct {
}

const (
	path    = "../log/"
	logfile = "application_error_%s_.log"
)

func WriteErrLog(appErr error) {
	currentTime := time.Now()
	YYYYMMDD := currentTime.Format("20060102")
	errlogfile := path + fmt.Sprintf(logfile, YYYYMMDD)

	file, err := os.OpenFile(errlogfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)

	log.Printf("%+v", appErr)
}
