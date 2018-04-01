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
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/arbor-dev/arbor/logger"
	"github.com/syndtr/goleveldb/leveldb"
)

type levelDBConnector struct {
	store *leveldb.DB
}

func newLevelDBConnector() *levelDBConnector {
	c := new(levelDBConnector)
	return c
}

func (c *levelDBConnector) open(location string) {
	var err error
	c.store, err = leveldb.OpenFile(location, nil)
	if err != nil {
		fmt.Println(err)
		logger.Log(logger.FATAL, err.Error())
	}
}

func (c *levelDBConnector) put(k []byte, v []byte) error {
	err := c.store.Put(k, v, nil)
	return err
}

func (c *levelDBConnector) get(k []byte) ([]byte, error) {
	v, err := c.store.Get(k, nil)
	return v, err
}

func (c *levelDBConnector) delete(v []byte) error {
	var err error
	iter := c.store.NewIterator(nil, nil)
	found := false
	for iter.Next() {
		currVal := iter.Value()
		if reflect.DeepEqual(currVal, v) {
			err = c.store.Delete(iter.Key(), nil)
			if err != nil {
				fmt.Println(err)
				logger.Log(logger.FATAL, err.Error())
			}
			found = true
		}
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		return err
	} else if found == false {
		return errors.New("No such value")
	}
	return nil
}

func (c *levelDBConnector) list() ([][]byte, error) {
	iter := c.store.NewIterator(nil, nil)
	values := make([][]byte, 0)
	for iter.Next() {
		v := iter.Value()
		c := make([]byte, len(v))
		copy(c, v)
		values = append(values, c)
	}
	iter.Release()
	err := iter.Error()
	return values, err
}

func (c *levelDBConnector) close() {
	err := c.store.Close()
	if err != nil {
		log.Fatal(err)
	}
}
