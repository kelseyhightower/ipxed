package web

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/kelseyhightower/httputils"
)

var (
	ErrCloudConfigAlreadyExists = errors.New("machine already exists")
)

type CloudConfig struct {
	Name   string `json:"name"`
	Config string `json:"config"`
}

func (c CloudConfig) Save() error {
	data, err := json.Marshal(&c)
	if err != nil {
		return err
	}
	resp, err := http.Post("http://localhost:8080/api/cloudconfigs", "application/json", bytes.NewReader(data))
	if err != nil {
		return err
	}
	switch resp.StatusCode {
	case http.StatusCreated:
		return nil
	case http.StatusConflict:
		return ErrProfileAlreadyExists
	}
	return nil
}

func DeleteCloudConfigByName(name string) error {
	req, err := http.NewRequest("DELETE", "http://localhost:8080/api/cloudconfigs/"+name, nil)
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

func GetCloudConfigByName(name string) (CloudConfig, error) {
	var c CloudConfig
	resp, err := http.Get("http://localhost:8080/api/cloudconfigs/" + name)
	if err != nil {
		return c, err
	}
	if resp.StatusCode != http.StatusOK {
		return c, errors.New("invalid response")
	}
	err = httputils.UnmarshalJSONBody(resp.Body, &c)
	if err != nil {
		return c, err
	}
	return c, nil
}

type CloudConfigsResponse struct {
	CloudConfigs []CloudConfig `json:"cloudconfigs"`
}

func GetCloudConfigs() ([]CloudConfig, error) {
	var response CloudConfigsResponse
	resp, err := http.Get("http://localhost:8080/api/cloudconfigs")
	if err != nil {
		return response.CloudConfigs, err
	}
	if resp.StatusCode != 200 {
		return response.CloudConfigs, errors.New(resp.Status)
	}
	err = httputils.UnmarshalJSONBody(resp.Body, &response)
	if err != nil {
		return response.CloudConfigs, err
	}
	return response.CloudConfigs, nil
}
