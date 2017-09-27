package logger

import (
	"log"
	"os"
)

//Sev is an enum for Logger Severity
type Sev int

const (
	//INFO Level Message
	INFO Sev = iota
	//WARN Level Message
	WARN Sev = iota
	//ERR Level Message
	ERR Sev = iota
	//FATAL Level Message
	FATAL Sev = iota
	//SPEC Level Message
	SPEC Sev = iota
)

//ColoredOutput controls if logs print in Color or Not
var ColoredOutput = true

//Log a messaage at a specific severity
func Log(sev Sev, msg string) {
	if ColoredOutput {
		switch sev {
		case WARN:
			log.Println("\x1b[33;1m[WARNING]: " + msg + "\x1b[0m")
		case ERR:
			log.Println("\x1b[31;1m[ERROR]: " + msg + "\x1b[0m")
		case SPEC:
			log.Println("\x1b[32;1m[ARBOR]: " + msg + "\x1b[0m")
		case FATAL:
			log.Println("\x1b[31;1m[FATAL]: " + msg + "\x1b[0m")
			os.Exit(1)
		default:
			log.Println("[INFO]: " + msg)
		}
		return
	}
	switch sev {
	case WARN:
		log.Println("[WARNING]: " + msg)
	case ERR:
		log.Println("[ERROR]: " + msg)
	case SPEC:
		log.Println("[ARBOR]: " + msg)
	case FATAL:
		log.Println("[FATAL]: " + msg)
		os.Exit(1)
	default:
		log.Println("[INFO]: " + msg)
	}
}
