package main

import (
	"fmt"
	"github.com/pingcap/tidb/kv"
	"github.com/pingcap/tidb/store/tikv"
)

type TiKV struct {
	store  kv.Storage
}

// Init initializes informations.
func NewTiKVServer(path string)(*TiKV, error) {
	driver := tikv.Driver{}
	store, err := driver.Open(fmt.Sprintf("tikv://%s?cluster=1",path ))
	if err != nil {
		return nil,err
	}
	return &TiKV{
		store: store,
	}, nil
}


func (t *TiKV) Get(key string)([]byte,error){
	txn, err := t.store.Begin()
	if err != nil {
		return err
	}
	return txn.Get(key)

}

func (t *TiKV) Set(key string, value []byte)error{
	txn, err := t.store.Begin()
	if err != nil {
		return err
	}
	txn.Set([]byte(key), value)
	err = txn.Commit()
	if err != nil {
		txn.Rollback()
		return  err
	}
	return nil
}