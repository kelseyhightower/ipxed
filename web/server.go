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

	// sshkeys
	r.HandleFunc("/sshkeys", IndexSSHKeyHandler).Methods("GET")
	r.HandleFunc("/sshkeys/create", CreateSSHKeyHandler).Methods("POST")
	r.HandleFunc("/sshkeys/new", NewSSHKeyHandler).Methods("GET")
	r.HandleFunc("/sshkeys/{name}", ShowSSHKeyHandler).Methods("GET")
	r.HandleFunc("/sshkeys/{name}/edit", EditSSHKeyHandler).Methods("GET")
	r.HandleFunc("/sshkeys/{name}/delete", DeleteSSHKeyHandler).Methods("GET")

	// cloudconfigs
	r.HandleFunc("/cloudconfigs", IndexCloudConfigHandler).Methods("GET")
	r.HandleFunc("/cloudconfigs/create", CreateCloudConfigHandler).Methods("POST")
	r.HandleFunc("/cloudconfigs/new", NewCloudConfigHandler).Methods("GET")
	r.HandleFunc("/cloudconfigs/{name}", ShowCloudConfigHandler).Methods("GET")
	r.HandleFunc("/cloudconfigs/{name}/edit", EditCloudConfigHandler).Methods("GET")
	r.HandleFunc("/cloudconfigs/{name}/delete", DeleteCloudConfigHandler).Methods("GET")

	return r
}
