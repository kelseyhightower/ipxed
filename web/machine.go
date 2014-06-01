package web

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

type Machine struct {
	Config     string
	Name       string
	MacAddress string
	Profile    string
	Serial     string
	UUID       string
}

func (m Machine) Save() {
	return
}

func GetMachineByName(name string) (Machine, error) {
	var m Machine
	return m, nil
}

type Response struct {
	Machines []Machine `json:"machines"`
}

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

func GetMachines() ([]Machine, error) {
	var response Response
	resp, err := http.Get("http://localhost:8080/api/machines")
	if err != nil {
		return response.Machines, err
	}
	if resp.StatusCode != 200 {
		return response.Machines, errors.New(resp.Status)
	}
	err = unmarshal(resp.Body, &response)
	if err != nil {
		return response.Machines, err
	}
	return response.Machines, nil
}
