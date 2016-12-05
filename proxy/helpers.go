/**
* Copyright © 2016, ACM@UIUC
*
* This file is part of the Groot Project.
*
* The Groot Project is open source software, released under the University of
* Illinois/NCSA Open Source License. You should have received a copy of
* this license in a file with the distribution.
**/

package proxy

import (
	"encoding/json"
	"errors"
	"github.com/acm-uiuc/groot/security"
	"net/http"
)

func contains(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func extract(a interface{}, b string) (val string, err error) {

	j, err := json.Marshal(a)
	if err != nil {
		return "", errors.New("key does not exist")
	}
	// a map container to decode the JSON structure into
	c := make(map[string]interface{})

	// unmarschal JSON
	e := json.Unmarshal(j, &c)

	// panic on error
	if e != nil {
		panic(e)
	}

	// a string slice to hold the keys
	k := make([]string, len(c))

	// iteration counter
	i := 0

	// copy c's keys into k
	for s, _ := range c {
		k[i] = s
		i++
	}

	if contains(b, k) {
		return c[b].(string), nil
	} else {
		return "", errors.New("key does not exist")
	}

}

func verifyAuthorization(r *http.Request) bool {
	authToken := r.Header.Get("Authorization")
	//IsAuthorizedClient Handles empty token
	auth, err := security.IsAuthorizedClient(authToken)
	if err != nil {
		return false
	}
	return auth
}
