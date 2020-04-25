package raft_h

import (
	"log"
	"testing"
)
import  "github.com/boltdb/bolt"
/*
  bolt 是 单机日志  的 lsm 结构  适合单机日志
 */

func TestBoltDb(t *testing.T)  {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		b,_ := tx.CreateBucket([]byte("hhh"))
		//b := tx.Bucket([]byte("hhh"))
		b.Put([]byte("hh"),[]byte("22222"))
		return nil
	})

}
