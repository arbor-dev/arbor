/**
* Copyright © 2017, ACM@UIUC
*
* This file is part of the Groot Project.
*
* The Groot Project is open source software, released under the University of
* Illinois/NCSA Open Source License. You should have received a copy of
* this license in a file with the distribution.
**/

package security

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/acm-uiuc/arbor/logger"
)

type accessLogger struct {
	accessLog *os.File
}

func newAccessLogger() *accessLogger {
	logger := new(accessLogger)
	return logger
}

func (l *accessLogger) open(accessLogLocation string) {
	_, err := os.Stat(accessLogLocation)

	if os.IsNotExist(err) {
		l.accessLog, err = os.Create(accessLogLocation)
		if err != nil {
			log.Fatal(err)
		}
	}

	l.accessLog, err = os.OpenFile(accessLogLocation, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
}

func (l *accessLogger) log(name string, token string) error {
	t := time.Now().Local()
	str := fmt.Sprintf("%s %s %s\n", t.Format("2006-01-02 15:04:05 +0800"), name, token)
	_, err := (*l.accessLog).WriteString(str)
	err = (*l.accessLog).Sync()
	logger.Log(logger.DEBUG, str)
	return err
}

func (l *accessLogger) close() {
	(*l.accessLog).Close()
}
