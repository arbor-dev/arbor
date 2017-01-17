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
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/kennygrant/sanitize"
)

func SanitizeRequest(r *http.Request) {

	content, _ := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	fmt.Println("Sanitizer before:" + string(content))
	fmt.Println(" after " + sanitize.HTML(string(content)))
	fmt.Println(" alt " + escapeString(string(content)))
	r.Body = ioutil.NopCloser(strings.NewReader(string(content)))
	r.ContentLength = int64(len(content))

}

var htmlEscaper = strings.NewReplacer(
	`&`, "&amp;",
	`<`, "&lt;",
	`>`, "&gt;",
)

// EscapeString escapes special characters like "<" to become "&lt;". It
// escapes only five such characters: <, >, &, ' and ".
// UnescapeString(EscapeString(s)) == s always holds, but the converse isn't
// always true.

func escapeString(s string) string {
	return htmlEscaper.Replace(s)
}
