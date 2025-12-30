package logger

import "log"

func Init(service string) {
	log.SetPrefix("[" + service + "] ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("logger initialized")
}
