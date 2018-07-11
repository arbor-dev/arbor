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
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/kennygrant/sanitize"
)

// Defines the maximum size for sanitized requests
const (
	MB = 1048576
	MaxSize = 16 * MB
)

func SanitizeRequest(r *http.Request) {
	if !enabled {
		return
	}
	content, _ := ioutil.ReadAll(io.LimitReader(r.Body, MaxSize))
	sanitizedHTML := sanitize.HTML(string(content))
	r.Body = ioutil.NopCloser(strings.NewReader(sanitizedHTML))
	r.ContentLength = int64(len(sanitizedHTML))
}
