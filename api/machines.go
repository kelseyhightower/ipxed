package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/httputils"
)

func CreateMachineHandler(w http.ResponseWriter, r *http.Request) {
	var m Machine
	err := httputils.UnmarshalJSONBody(r.Body, &m)
	if err != nil {
		httputils.JSONError(w, err.Error(), 500)
	}
	err = m.Save()
	if err != nil {
		httputils.JSONError(w, err.Error(), 500)
	}
	httputils.SetLocationHeader(w, "http://localhost:8080/machines/"+m.Name)
	w.WriteHeader(http.StatusCreated)
}

func GetMachineHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	m, err := GetMachineByName(vars["name"])
	if err == ErrMachineNotFound {
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
	data, err := json.MarshalIndent(&response, "", "  ")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
	}
	w.Write(data)
}

func UpdateMachineHandler(w http.ResponseWriter, r *http.Request) {}
func DeleteMachineHandler(w http.ResponseWriter, r *http.Request) {}
