package web

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/kelseyhightower/httputils"
)

var (
	ErrMachineAlreadyExists = errors.New("machine already exists")
)

type Machine struct {
	Name       string `json:"name"`
	MacAddress string `json:"macaddress"`
	Profile    string `json:"profile"`
}

func (m Machine) Save() error {
	data, err := json.Marshal(&m)
	if err != nil {
		return err
	}
	resp, err := http.Post("http://localhost:8080/api/machines", "application/json", bytes.NewReader(data))
	if err != nil {
		return err
	}
	switch resp.StatusCode {
	case http.StatusCreated:
		return nil
	case http.StatusConflict:
		return ErrMachineAlreadyExists
	}
	return nil
}

func DeleteMachineByName(name string) error {
	req, err := http.NewRequest("DELETE", "http://localhost:8080/api/machines/"+name, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return errors.New("invalid response")
	}
	return nil
}

func GetMachineByName(name string) (Machine, error) {
	var m Machine
	resp, err := http.Get("http://localhost:8080/api/machines/" + name)
	if err != nil {
		return m, err
	}
	if resp.StatusCode != http.StatusOK {
		return m, errors.New("invalid response")
	}
	err = httputils.UnmarshalJSONBody(resp.Body, &m)
	if err != nil {
		return m, err
	}
	return m, nil
}

type Response struct {
	Machines []Machine `json:"machines"`
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
	err = httputils.UnmarshalJSONBody(resp.Body, &response)
	if err != nil {
		return response.Machines, err
	}
	return response.Machines, nil
}
