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
)

// Add a Authroized API Client to Arbor
// Will return an API key on successful addition
// Will return DB error if there is an issue
func AddClient(name string) (string, error) {
	token, err := generateRandomString(32)
	if err != nil {
		return "", err
	}
	err = clientRegistry.put([]byte(token), []byte(name))
	if err != nil {
		return "", err
	}
	return token, nil
}

// Verify if a key provided by a client is vaild
func IsAuthorizedClient(token string) (bool, error) {
	if !enabled {
		return true, nil
	}
	name, err := clientRegistry.get([]byte(token))
	if err != nil {
		return false, err
	}
	nameStr := string(name)
	if nameStr == "" {
		return false, fmt.Errorf("Not a valid token")
	}
	accessLog.log(nameStr, token)
	return true, nil
}

func RemoveClient(token string) error {
	return nil
}
