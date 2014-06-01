package api

import (
	"log"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

func init() {
	var err error
	db, err = bolt.Open("/tmp/ipxeweb.bolt", 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("machines"))
		if err != nil {
			log.Fatal(err.Error())
		}
		return nil
	})
	if err != nil {
		log.Fatal(err.Error())
	}
}
