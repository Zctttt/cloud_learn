package debug_log

import (
	"log"
	"os"
)

func Info(info string) {
	logFile, err := os.Create("log")
	if err != nil {
		log.Fatalln("open file error")
	}
	debugLog := log.New(logFile, "[info]", log.Ldate)
	debugLog.Println(info)
}
