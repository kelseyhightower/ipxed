package api

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

var (
	ErrNotFound = errors.New("not found")
)

func init() {
	var err error
	db, err = bolt.Open("/tmp/ipxeweb.bolt", 0644, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = db.Update(func(tx *bolt.Tx) error {
		buckets := []string{
			"machines",
			"profiles",
			"sshkeys",
			"cloudconfigs",
		}
		for _, b := range buckets {
			_, err := tx.CreateBucketIfNotExists([]byte(b))
			if err != nil {
				log.Fatal(err.Error())
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err.Error())
	}
}

func DeleteFromBucket(bucket, name string) error {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		b.Delete([]byte(name))
		return nil
	})
	return nil
}

func GetFromBucket(bucket, name string, v interface{}) error {
	var data []byte
	err := db.View(func(tx *bolt.Tx) error {
		data = tx.Bucket([]byte(bucket)).Get([]byte(name))
		return nil
	})
	if err != nil {
		return err
	}
	if len(data) == 0 {
		return ErrNotFound
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	return nil
}

func PutToBucket(bucket, name string, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		b.Put([]byte(name), data)
		return nil
	})
	return nil
}

func GetAllFromBucket(bucket string, f func(k, v []byte) error) error {
	return db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		// Iterate over items in sorted key order.
		return b.ForEach(f)
	})
}
