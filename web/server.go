package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Handler() http.Handler {
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	r.HandleFunc("/machines/create/", CreateMachineHandler).Methods("GET")
	r.HandleFunc("/machines/{name}", ShowMachineHandler).Methods("GET")
	r.HandleFunc("/", ListMachinesHandler).Methods("GET")
	r.HandleFunc("/machines/", ListMachinesHandler).Methods("GET")
	r.HandleFunc("/machines/edit/{name}", EditMachineHandler).Methods("GET")
	r.HandleFunc("/machines/save/", SaveMachineHandler).Methods("POST")
	return r
}
