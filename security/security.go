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

//Default location for access log
var AccessLogLocation string = "log/access.log"

//Default location for client db
var ClientRegistryLocation string = "clients.db"

var accessLog *accessLogger
var clientRegistry *levelDBConnector

var enabled = false

func Init() {
	enabled = true
	clientRegistry = newLevelDBConnector()
	accessLog = newAccessLogger()
	clientRegistry.open(ClientRegistryLocation)
	accessLog.open(AccessLogLocation)
}

func IsEnabled() bool {
	return enabled
}

func Shutdown() {
	clientRegistry.close()
	accessLog.close()
	enabled = false
}
