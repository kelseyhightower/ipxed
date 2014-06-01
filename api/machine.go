package api

import (
	"encoding/json"
	"errors"

	"github.com/boltdb/bolt"
)

var (
	ErrMachineNotFound = errors.New("machine not found")
)

type Machine struct {
	Name       string `json:"name"`
	MacAddress string `json:"macaddress"`
	Profile    string `json:"profile"`
}

func (m Machine) Save() error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("machines"))
		b.Put([]byte(m.Name), data)
		return nil
	})
	return nil
}

func (m Machine) Delete() error {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("machines"))
		b.Delete([]byte(m.Name))
		return nil
	})
	return nil
}

func GetMachines() ([]Machine, error) {
	machines := make([]Machine, 0)
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("machines"))
		// Iterate over items in sorted key order.
		b.ForEach(func(k, v []byte) error {
			var m Machine
			err := json.Unmarshal(v, &m)
			if err != nil {
				return err
			}
			machines = append(machines, m)
			return nil
		})
		return nil
	})
	return machines, nil
}

func GetMachineByName(name string) (Machine, error) {
	var m Machine
	var data []byte
	err := db.View(func(tx *bolt.Tx) error {
		data = tx.Bucket([]byte("machines")).Get([]byte(name))
		return nil
	})
	if err != nil {
		return m, err
	}
	if len(data) == 0 {
		return m, ErrMachineNotFound
	}
	err = json.Unmarshal(data, &m)
	if err != nil {
		return m, err
	}
	return m, nil
}
