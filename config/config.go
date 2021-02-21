package config

import (
	"log"
	"time"
)

var err error
var timeStart time.Time

func InitApp() {
	setLog()
	loadEnv()
	startTime()
}

func setLog() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

func startTime() {
	timeStart = time.Now()
}

func GetTime() time.Time {
	return timeStart
}
