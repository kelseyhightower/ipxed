package main

import (
	"log"
	"net/http"

	"github.com/boltdb/bolt"
	"github.com/gorilla/mux"
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

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	r.HandleFunc("/machines/create/", CreateMachineHandler).Methods("GET")
	r.HandleFunc("/machines/{name}", ShowMachineHandler).Methods("GET")
	r.HandleFunc("/", ListMachinesHandler).Methods("GET")
	r.HandleFunc("/machines/", ListMachinesHandler).Methods("GET")
	r.HandleFunc("/machines/edit/{name}", EditMachineHandler).Methods("GET")
	r.HandleFunc("/machines/save/", SaveMachineHandler).Methods("POST")
	http.Handle("/", r)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
