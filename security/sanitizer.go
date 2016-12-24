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
	//"html"
	"net/http"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

func SanitizeRequest(r *http.Request) {
	
	content, _ := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	fmt.Println("Sanitizer " + string(content));
	r.Body = ioutil.NopCloser(strings.NewReader(string(content)))
	r.ContentLength = int64(len(content))	

}




 