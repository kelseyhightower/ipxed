package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Handler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/api/machines", CreateMachineHandler).Methods("POST")
	r.HandleFunc("/api/machines", ListMachinesHandler).Methods("GET")
	return r
}
