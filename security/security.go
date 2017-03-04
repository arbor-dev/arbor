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

var AccessLogLocation string = "log/access.log"
var ClientRegistryLocation string = "clients.db"
var enabled = false

func Init() {
	enabled = true
	storeOpen()
	logOpen(AccessLogLocation)
}

func Shutdown() {
	storeClose()
	logClose()
	enabled = false
}
