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
	"log"

	"github.com/acm-uiuc/arbor/logger"
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

func (c *levelDBConnector) delete(k []byte) error {
	err := c.store.Delete(k, nil)
	return err
}

func (c *levelDBConnector) close() {
	err := c.store.Close()
	if err != nil {
		log.Fatal(err)
	}
}
