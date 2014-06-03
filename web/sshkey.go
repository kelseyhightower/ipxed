package web

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/kelseyhightower/httputils"
)

var (
	ErrSSHKeyAlreadyExists = errors.New("machine already exists")
)

type SSHKey struct {
	Name        string `json:"name"`
	Fingerprint string `json:"fingerprint"`
	Key         string `json:"key"`
}

func (s SSHKey) Save() error {
	data, err := json.Marshal(&s)
	if err != nil {
		return err
	}
	resp, err := http.Post("http://localhost:8080/api/sshkeys", "application/json", bytes.NewReader(data))
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

func DeleteSSHKeyByName(name string) error {
	req, err := http.NewRequest("DELETE", "http://localhost:8080/api/sshkeys/"+name, nil)
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

func GetSSHKeyByName(name string) (SSHKey, error) {
	var s SSHKey
	resp, err := http.Get("http://localhost:8080/api/sshkeys/" + name)
	if err != nil {
		return s, err
	}
	if resp.StatusCode != http.StatusOK {
		return s, errors.New("invalid response")
	}
	err = httputils.UnmarshalJSONBody(resp.Body, &s)
	if err != nil {
		return s, err
	}
	return s, nil
}

type SSHKeysResponse struct {
	SSHKeys []SSHKey `json:"sshkeys"`
}

func GetSSHKeys() ([]SSHKey, error) {
	var response SSHKeysResponse
	resp, err := http.Get("http://localhost:8080/api/sshkeys")
	if err != nil {
		return response.SSHKeys, err
	}
	if resp.StatusCode != 200 {
		return response.SSHKeys, errors.New(resp.Status)
	}
	err = httputils.UnmarshalJSONBody(resp.Body, &response)
	if err != nil {
		return response.SSHKeys, err
	}
	return response.SSHKeys, nil
}
