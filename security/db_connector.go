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
	"log"
	"time"

	"github.com/boltdb/bolt"
)

func storeOpen() {
	var err error
	clientRegistryStore, err = bolt.Open(clientRegistryLocation, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
}

func put(k []byte, v []byte) error {
	err := clientRegistryStore.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("clients"))
		if err != nil {
			return err
		}
		err = bucket.Put(k, v)
		if err != nil {
			return err
		}
		return err
	})
	return err
}

func get(k []byte) ([]byte, error) {
	var v []byte
	err := clientRegistryStore.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("clients"))
		if b == nil {
			return fmt.Errorf("Bucket client not found")
		}
		v = b.Get(k)
		return nil
	})
	return v, err
}

func storeClose() {
	clientRegistryStore.Close()
}
