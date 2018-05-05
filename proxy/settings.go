/**
* Copyright © 2017, ACM@UIUC
*
* This file is part of the Groot Project.
*
* The Groot Project is open source software, released under the University of
* Illinois/NCSA Open Source License. You should have received a copy of
* this license in a file with the distribution.
**/

package proxy

// JSONHeader is the default header for the JSON content-type
var JSONHeader = "application/json; charset=UTF-8"

// TEXTHeader is the default header for the plain text content-type
var TEXTHeader = "text/plain; charset=utf-8"

// HTMLHeader is the default header for the HTML content-type
var HTMLHeader = "text/html;charset=utf-8"

// XMLHeader is the default header for the XML content-type
var XMLHeader = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>"

// Timeout is the default request timeout
var Timeout int64 = 10

// AccessControlPolicy is the default Access control policy
var AccessControlPolicy = "*"

// Client Authorization Token Field
var ClientAuthorizationHeaderField = "Authorization"
