package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Handler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/api/machines", ListMachinesHandler).Methods("GET")
	r.HandleFunc("/api/machines", CreateMachineHandler).Methods("POST")
	r.HandleFunc("/api/machines/{name}", GetMachineHandler).Methods("GET")
	r.HandleFunc("/api/machines/{name}", UpdateMachineHandler).Methods("PUT")
	r.HandleFunc("/api/machines/{name}", DeleteMachineHandler).Methods("DELETE")
	return r
}
