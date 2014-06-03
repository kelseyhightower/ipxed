package web

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// CreateSSHKeyHandler creates a new sshkey.
func CreateSSHKeyHandler(w http.ResponseWriter, r *http.Request) {
	s := SSHKey{}
	s.Name = r.PostFormValue("name")
	s.Key = r.PostFormValue("sshkey")
	s.Save()
	http.Redirect(w, r, "/sshkeys/"+s.Name, http.StatusMovedPermanently)
}

// DeleteSSHKeyHandler deletes a specific sshkey.
func DeleteSSHKeyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	DeleteSSHKeyByName(vars["name"])
	http.Redirect(w, r, "/sshkeys", http.StatusMovedPermanently)
}

// EditSSHKeyHandler displays an HTML form for editing a sshkey.
func EditSSHKeyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	s, err := GetSSHKeyByName(name)
	if err != nil {
		log.Println(err.Error())
	}
	p := &Page{
		Data:  s,
		Title: "Edit " + name,
	}
	renderTemplate(w, "templates/sshkeys/edit.html", p)
}

// IndexSSHKeyHandler displays a list of all sshkeys.
func IndexSSHKeyHandler(w http.ResponseWriter, r *http.Request) {
	keys, err := GetSSHKeys()
	if err != nil {
		log.Println(err.Error())
	}
	p := &Page{
		Data:  keys,
		Title: "SSH Keys",
	}
	renderTemplate(w, "templates/sshkeys/index.html", p)
}

// NewSSHKeyHandler displays an HTML form for creating a sshkey.
func NewSSHKeyHandler(w http.ResponseWriter, r *http.Request) {
	p := &Page{Title: "Add SSH Key"}
	renderTemplate(w, "templates/sshkeys/create.html", p)
}

// ShowSSHKeyHandler displays a specific sshkey.
func ShowSSHKeyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	s, err := GetSSHKeyByName(name)
	if err != nil {
		log.Println(err.Error())
	}
	p := &Page{
		Data:  s,
		Title: s.Name,
	}
	renderTemplate(w, "templates/sshkeys/show.html", p)
}
