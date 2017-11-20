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
	namestr := string(name)
	if namestr == "" {
		return false, fmt.Errorf("Not a valid token")
	}

	accessLog.log(namestr, token)
	return true, nil
}

func DeleteClient(name string) error {
	return clientRegistry.delete([]byte(name))
}

func ListClients() ([]string, error) {
	clients, err := clientRegistry.list()
	if err != nil {
		return []string{}, err
	}
	names := make([]string, len(clients))
	for i := 0; i < len(clients); i++ {
		names[i] = string(clients[i])
	}
	return names, nil
}
