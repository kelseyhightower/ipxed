package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Handler() http.Handler {
	r := mux.NewRouter()
	// machines
	r.HandleFunc("/api/machines", ListMachinesHandler).Methods("GET")
	r.HandleFunc("/api/machines", CreateMachineHandler).Methods("POST")
	r.HandleFunc("/api/machines/{name}", GetMachineHandler).Methods("GET")
	r.HandleFunc("/api/machines/{name}", UpdateMachineHandler).Methods("PUT")
	r.HandleFunc("/api/machines/{name}", DeleteMachineHandler).Methods("DELETE")

	// profiles
	r.HandleFunc("/api/profiles", ListProfilesHandler).Methods("GET")
	r.HandleFunc("/api/profiles", CreateProfileHandler).Methods("POST")
	r.HandleFunc("/api/profiles/{name}", GetProfileHandler).Methods("GET")
	r.HandleFunc("/api/profiles/{name}", DeleteProfileHandler).Methods("DELETE")

	// sshkeys
	r.HandleFunc("/api/sshkeys", ListSSHKeyHandler).Methods("GET")
	r.HandleFunc("/api/sshkeys", CreateSSHKeyHandler).Methods("POST")
	r.HandleFunc("/api/sshkeys/{name}", GetSSHKeyHandler).Methods("GET")
	r.HandleFunc("/api/sshkeys/{name}", DeleteSSHKeyHandler).Methods("DELETE")


	return r
}
