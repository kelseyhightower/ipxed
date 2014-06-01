package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Handler() http.Handler {
	r := mux.NewRouter()
	// Static assets
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	// machines endpoint
	r.HandleFunc("/", ListMachinesHandler).Methods("GET")
	r.HandleFunc("/machines/", ListMachinesHandler).Methods("GET")
	r.HandleFunc("/machines/{name}", ShowMachineHandler).Methods("GET")
	r.HandleFunc("/machines/create/", CreateMachineHandler).Methods("GET")
	r.HandleFunc("/machines/delete/{name}", DeleteMachineHandler).Methods("POST")
	r.HandleFunc("/machines/edit/{name}", EditMachineHandler).Methods("POST")
	r.HandleFunc("/machines/save/", SaveMachineHandler).Methods("POST")
	return r
}
