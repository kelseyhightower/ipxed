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

	// machines
	r.HandleFunc("/machines", IndexMachineHandler).Methods("GET")
	r.HandleFunc("/machines/create", CreateMachineHandler).Methods("POST")
	r.HandleFunc("/machines/new", NewMachineHandler).Methods("GET")
	r.HandleFunc("/machines/{name}", ShowMachineHandler).Methods("GET")
	r.HandleFunc("/machines/{name}/edit", EditMachineHandler).Methods("GET")
	r.HandleFunc("/machines/{name}/delete", DeleteMachineHandler).Methods("GET")

	// profiles
	r.HandleFunc("/profiles", IndexProfileHandler).Methods("GET")
	r.HandleFunc("/profiles/create", CreateProfileHandler).Methods("POST")
	r.HandleFunc("/profiles/new", NewProfileHandler).Methods("GET")
	r.HandleFunc("/profiles/{name}", ShowProfileHandler).Methods("GET")
	r.HandleFunc("/profiles/{name}/edit", EditProfileHandler).Methods("GET")
	r.HandleFunc("/profiles/{name}/delete", DeleteProfileHandler).Methods("GET")

	return r
}
