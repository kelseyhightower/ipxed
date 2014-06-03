package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// CreateCloudConfigsHandler creates a new cloudconfig.
func CreateCloudConfigHandler(w http.ResponseWriter, r *http.Request) {
	c := CloudConfig{}
	c.Name = r.PostFormValue("name")
	c.Config = r.PostFormValue("config")
	c.Save()
	http.Redirect(w, r, "/cloudconfigs/"+c.Name, http.StatusMovedPermanently)
}

// DeleteCloudConfigHandler deletes a specific cloudconfig.
func DeleteCloudConfigHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	DeleteCloudConfigByName(vars["name"])
	http.Redirect(w, r, "/cloudconfigs", http.StatusMovedPermanently)
}

// EditCloudConfigHandler displays an HTML form for editing a cloudconfig.
func EditCloudConfigHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	c, err := GetCloudConfigByName(vars["name"])
	if err != nil {
		log.Println(err.Error())
	}
	p := &Page{
		Data:  c,
		Title: fmt.Sprintf("Cloud Config - %s", c.Name),
	}
	renderTemplate(w, "templates/cloudconfigs/edit.html", p)
}

// IndexCloudConfigHandler displays a list of all cloudconfigs.
func IndexCloudConfigHandler(w http.ResponseWriter, r *http.Request) {
	cloudconfigs, err := GetCloudConfigs()
	if err != nil {
		log.Println(err.Error())
	}
	p := &Page{
		Data:  cloudconfigs,
		Title: "Cloud Configs",
	}
	renderTemplate(w, "templates/cloudconfigs/index.html", p)
}

// NewCloudConfigHandler displays an HTML form for creating a cloudconfig.
func NewCloudConfigHandler(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title: "Add Cloud Config",
	}
	renderTemplate(w, "templates/cloudconfigs/create.html", p)
}

// ShowCloudConfigHandler displays a specific cloudconfig.
func ShowCloudConfigHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	c, err := GetCloudConfigByName(vars["name"])
	if err != nil {
		log.Println(err.Error())
	}
	p := &Page{
		Data:  c,
		Title: fmt.Sprintf("Cloud Config - %s", c.Name),
	}
	renderTemplate(w, "templates/cloudconfigs/show.html", p)
}
