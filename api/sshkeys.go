package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/httputils"
)

func CreateSSHKeyHandler(w http.ResponseWriter, r *http.Request) {
	var s SSHKey
	if err := httputils.UnmarshalJSONBody(r.Body, &s); err != nil {
		httputils.JSONError(w, err.Error(), 500)
	}
	if err := s.Save(); err != nil {
		httputils.JSONError(w, err.Error(), 500)
	}
	httputils.SetLocationHeader(w, "http://localhost:8080/sshkeys/"+s.Name)
	w.WriteHeader(http.StatusCreated)
}

func GetSSHKeyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	s, err := GetSSHKeyByName(vars["name"])
	if err == ErrNotFound {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		httputils.JSONError(w, err.Error(), 500)
		return
	}
	httputils.JSONWrite(w, s, 200)
}

type SSHKeysResponse struct {
	SSHKeys []SSHKey `json:"sshkeys"`
}

func ListSSHKeyHandler(w http.ResponseWriter, r *http.Request) {
	var response SSHKeysResponse
	sshkeys, err := GetSSHKeys()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
	}
	response.SSHKeys = sshkeys
	httputils.JSONWrite(w, response, 200)
}

func DeleteSSHKeyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := DeleteSSHKeyByName(vars["name"])
	if err != nil {
		httputils.JSONError(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
