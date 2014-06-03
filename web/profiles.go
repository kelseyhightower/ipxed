package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// CreateProfileHandler creates a new profile.
func CreateProfileHandler(w http.ResponseWriter, r *http.Request) {
	profile := Profile{
		Name:            r.PostFormValue("name"),
		CloudConfig:     r.PostFormValue("cloud_config"),
		Console:         r.PostFormValue("console"),
		CoreosAutologin: r.PostFormValue("coreos_autologin"),
		RootFstype:      r.PostFormValue("rootfstype"),
		Root:            r.PostFormValue("root"),
		SSHKey:          r.PostFormValue("sshkey"),
		Version:         r.PostFormValue("version"),
	}
	profile.Save()
	http.Redirect(w, r, "/profiles/"+profile.Name, http.StatusMovedPermanently)
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
	profile, err := GetProfileByName(vars["name"])
	if err != nil {
		log.Println(err.Error())
	}
	page := &Page{
		Data:  profile,
		Title: fmt.Sprintf("Profiles - %s", profile.Name),
	}
	renderTemplate(w, "templates/profiles/edit.html", page)
}

// IndexProfileHandler displays a list of all profiles.
func IndexProfileHandler(w http.ResponseWriter, r *http.Request) {
	profiles, err := GetProfiles()
	if err != nil {
		log.Println(err.Error())
	}
	page := &Page{
		Data:  profiles,
		Title: "Profiles",
	}
	renderTemplate(w, "templates/profiles/index.html", page)
}

// NewProfileHandler displays an HTML form for creating a profile.
func NewProfileHandler(w http.ResponseWriter, r *http.Request) {
	page := &Page{
		Title: "Create a Profile",
	}
	renderTemplate(w, "templates/profiles/create.html", page)
}

// ShowProfileHandler displays a specific profile.
func ShowProfileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	profile, err := GetProfileByName(vars["name"])
	if err != nil {
		log.Println(err.Error())
	}
	page := &Page{
		Data:  profile,
		Title: fmt.Sprintf("Profiles - %s", profile.Name),
	}
	renderTemplate(w, "templates/profiles/show.html", page)
}
