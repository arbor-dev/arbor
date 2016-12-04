/**
* Copyright © 2016, ACM@UIUC
*
* This file is part of the Groot Project.  
* 
* The Groot Project is open source software, released under the University of
* Illinois/NCSA Open Source License. You should have received a copy of
* this license in a file with the distribution.
**/

package security

import (
	"github.com/acm-uiuc/groot/config"
	"os"
	"time"
	"fmt"
	"log"
)

const AccessLogLocation string = config.AccessLogLocation
var AccessLog *os.File 

func logInit() {
	_, err := os.Stat(AccessLogLocation)

	if os.IsNotExist(err) {
		AccessLog, err = os.Create(AccessLogLocation) 
		if err != nil {
			log.Fatal(err)
		}
	}

	AccessLog, err = os.OpenFile(AccessLogLocation, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
}

func appendLog(name string, token string) (error) {
	t := time.Now().Local()
	str := fmt.Sprintf("%s %s %s", t.Format("2006-01-02 15:04:05 +0800"), name, token)
	_, err := (*AccessLog).WriteString(str)
	err = (*AccessLog).Sync()
	return err
} 

func logClose() {
	(*AccessLog).Close()
}
