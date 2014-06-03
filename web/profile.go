package web

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/kelseyhightower/httputils"
)

var (
	ErrProfileAlreadyExists = errors.New("machine already exists")
)

type Profile struct {
	Name            string `json:"name"`
	CloudConfig     string `json:"cloud_config"`
	Console         string `json:"console"`
	CoreosAutologin string `json:"coreos_autologin"`
	RootFstype      string `json:"rootfstype"`
	Root            string `json:"root"`
	SSHKey          string `json:"sshkey"`
	Version         string `json:"version"`
}

func (p Profile) Save() error {
	data, err := json.Marshal(&p)
	if err != nil {
		return err
	}
	resp, err := http.Post("http://localhost:8080/api/profiles", "application/json", bytes.NewReader(data))
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

func DeleteProfileByName(name string) error {
	req, err := http.NewRequest("DELETE", "http://localhost:8080/api/profiles/"+name, nil)
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

func GetProfileByName(name string) (Profile, error) {
	var p Profile
	resp, err := http.Get("http://localhost:8080/api/profiles/" + name)
	if err != nil {
		return p, err
	}
	if resp.StatusCode != http.StatusOK {
		return p, errors.New("invalid response")
	}
	err = httputils.UnmarshalJSONBody(resp.Body, &p)
	if err != nil {
		return p, err
	}
	return p, nil
}

type ProfilesResponse struct {
	Profiles []Profile `json:"profiles"`
}

func GetProfiles() ([]Profile, error) {
	var response ProfilesResponse
	resp, err := http.Get("http://localhost:8080/api/profiles")
	if err != nil {
		return response.Profiles, err
	}
	if resp.StatusCode != 200 {
		return response.Profiles, errors.New(resp.Status)
	}
	err = httputils.UnmarshalJSONBody(resp.Body, &response)
	if err != nil {
		return response.Profiles, err
	}
	return response.Profiles, nil
}
