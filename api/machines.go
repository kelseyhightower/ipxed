package api

import (
	"log"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func unmarshal(r io.Reader, v interface{}) error {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		return err
	}
	return nil
}

func CreateMachineHandler(w http.ResponseWriter, r *http.Request) {
	var m Machine
	err := unmarshal(r.Body, &m)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	err = m.Save()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

type Response struct {
	Machines []Machine `json:"machines"`
}

func ListMachinesHandler(w http.ResponseWriter, r *http.Request) {
	var response Response
	machines, err := GetMachines()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
	}
	response.Machines = machines
	data, err := json.MarshalIndent(&response, "", "  ")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
	}
	w.Write(data)
}
