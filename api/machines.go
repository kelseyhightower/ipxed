package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/httputils"
)

func CreateMachineHandler(w http.ResponseWriter, r *http.Request) {
	var m Machine
	if err := httputils.UnmarshalJSONBody(r.Body, &m); err != nil {
		httputils.JSONError(w, err.Error(), 500)
	}
	if err := m.Save(); err != nil {
		httputils.JSONError(w, err.Error(), 500)
	}
	httputils.SetLocationHeader(w, "http://localhost:8080/machines/"+m.Name)
	w.WriteHeader(http.StatusCreated)
}

func GetMachineHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	m, err := GetMachineByName(vars["name"])
	if err == ErrNotFound {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		httputils.JSONError(w, err.Error(), 500)
		return
	}
	httputils.JSONWrite(w, m, 200)
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
	httputils.JSONWrite(w, response, 200)
}

func UpdateMachineHandler(w http.ResponseWriter, r *http.Request) {}

func DeleteMachineHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := DeleteMachineByName(vars["name"])
	if err != nil {
		httputils.JSONError(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
