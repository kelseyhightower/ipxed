package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Handler() http.Handler {
	r := mux.NewRouter()
	// Static assets.
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	r.HandleFunc("/", IndexMachineHandler).Methods("GET")

	// Machine endpoint.
	r.HandleFunc("/machines", IndexMachineHandler).Methods("GET")
	r.HandleFunc("/machines/create", CreateMachineHandler).Methods("POST")
	r.HandleFunc("/machines/new", NewMachineHandler).Methods("GET")
	r.HandleFunc("/machines/{name}", ShowMachineHandler).Methods("GET")
	r.HandleFunc("/machines/{name}/edit", EditMachineHandler).Methods("GET")
	r.HandleFunc("/machines/{name}/delete", DeleteMachineHandler).Methods("GET")
	return r
}
