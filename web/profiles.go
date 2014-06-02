package web

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// CreateProfileHandler creates a new profile.
func CreateProfileHandler(w http.ResponseWriter, r *http.Request) {
	p := Profile{}
	p.Name = r.PostFormValue("name")
	p.CloudConfig = r.PostFormValue("cloud_config")
	p.Console = r.PostFormValue("console")
	p.CoreosAutologin = r.PostFormValue("coreos_autologin")
	p.RootFstype = r.PostFormValue("rootfstype")
	p.Root = r.PostFormValue("root")
	p.SSHKey = r.PostFormValue("sshkey")
	p.Version = r.PostFormValue("version")
	p.Save()
	http.Redirect(w, r, "/profiles/"+p.Name, http.StatusMovedPermanently)
}

// DeleteProfileHandler deletes a specific profile.
func DeleteProfileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	DeleteProfileByName(vars["name"])
	http.Redirect(w, r, "/profiles", http.StatusMovedPermanently)
}

// EditProfileHandler displays an HTML form for editing a profile.
func EditProfileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	pf, err := GetProfileByName(name)
	if err != nil {
		log.Println(err.Error())
	}
	p := &Page{
		Data:  pf,
		Title: "Edit " + name,
	}
	renderTemplate(w, "templates/profiles/edit.html", p)
}

// IndexProfileHandler displays a list of all profiles.
func IndexProfileHandler(w http.ResponseWriter, r *http.Request) {
	profiles, err := GetProfiles()
	if err != nil {
		log.Println(err.Error())
	}
	p := &Page{
		Data:  profiles,
		Title: "Profiles",
	}
	renderTemplate(w, "templates/profiles/index.html", p)
}

// NewProfileHandler displays an HTML form for creating a profile.
func NewProfileHandler(w http.ResponseWriter, r *http.Request) {
	p := &Page{ Title: "Create Profile"}
	renderTemplate(w, "templates/profiles/create.html", p)
}

// ShowProfileHandler displays a specific profile.
func ShowProfileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	pf, err := GetProfileByName(name)
	if err != nil {
		log.Println(err.Error())
	}
	p := &Page{
		Data:  pf,
		Title: pf.Name,
	}
	renderTemplate(w, "templates/profiles/show.html", p)
}
