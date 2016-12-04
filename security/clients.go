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
	"log"
	"fmt"
	"time"
	"github.com/boltdb/bolt"
	"github.com/acm-uiuc/groot/config"
)

const ClientRegistryLocation string = config.ClientRegistryLocation
var ClientRegistryStore *bolt.DB

func storeInit() {
	ClientRegistryStore, err := bolt.Open(ClientRegistryLocation, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer ClientRegistryStore.Close()
} 

func AddClient(name string) (string, error) {
	token, err := generateRandomString(32)
	if err  != nil {
		return "", err
	}
	err = put([]byte(token), []byte(name))
	if err != nil {
		return "", err
	}
	return token, nil 
}

func IsAuthorizedClient(token string) (bool, error) {
	name, err :=  get([]byte(token))
	if err != nil {
		return false, err
	}
	name_str := string(name)
	if name_str == "" {
		return false, fmt.Errorf("Not a valid token")
	} else {
		appendLog(name_str, token)
		return true, nil
	}
}

func RemoveClient(token string) (error) {
	return nil
}
