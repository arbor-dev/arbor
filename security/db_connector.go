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
	"github.com/boltdb/bolt"
)

func put(k []byte, v []byte) (error) {
	err := ClientRegistryStore.Update(func(tx *bolt.Tx) (error) {
		b, err := tx.CreateBucketIfNotExists([]byte("clients"))
		if err != nil {
			return err
		}
		err = b.Put(k, v)
		if err != nil {
			return err
		}
		return err
	})
	return err
}

func get(k []byte) ([]byte, error) {
	var v []byte
	err := ClientRegistryStore.View(func(tx *bolt.Tx) (error) {
		b := tx.Bucket([]byte("clients"))
		if b == nil {
			return fmt.Errorf("Bucket client not found")
		}
		v = b.Get(k)
		return nil
	})
	return v, err
}
