/**
* Copyright © 2017, ACM@UIUC
*
* This file is part of the Groot Project.
*
* The Groot Project is open source software, released under the University of
* Illinois/NCSA Open Source License. You should have received a copy of
* this license in a file with the distribution.
**/

package logger

import (
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

//Sev is an enum for Logger Severity
type Sev int

const (
	//DEBUG level Message
	DEBUG Sev = 3
	//INFO Level Message
	INFO Sev = 2
	//WARN Level Message
	WARN Sev = 1
	//ERR Level Message
	ERR Sev = 0
	//SPEC Level Message
	SPEC Sev = -1
	//FATAL Level Message
	FATAL Sev = -2
)

//ColoredOutput controls if logs print in Color or Not
var ColoredOutput = true

//LogLevel controls the verboseness of logging
var LogLevel = DEBUG

//Log a messaage at a specific severity
func Log(sev Sev, msg string) {
	if !(LogLevel >= sev) && !(sev == FATAL) {
		return
	}
	if ColoredOutput {
		switch sev {
		case DEBUG:
			log.Println("[DEBUG]: " + msg)
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
	case DEBUG:
		log.Println("[DEBUG]: " + msg)
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

//LogReq is a helper to log requests
func LogReq(sev Sev, req *http.Request) {
	rDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		Log(ERR, err.Error())
		return
	}
	Log(sev, string("Request:\n\n")+string(rDump))
}

func LogResp(sev Sev, resp *http.Response) {
	rDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		Log(ERR, err.Error())
		return
	}
	Log(sev, string("Response:\n\n")+string(rDump)+string('\n'))
}
