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
)

var enabled = false

func Init() {
	enabled = true
	storeInit()
	logInit()
}

func Close() {
	storeClose()
	logClose()
} 