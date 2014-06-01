package web

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func SaveMachineHandler(w http.ResponseWriter, r *http.Request) {
	m := Machine{}
	m.Name = r.PostFormValue("name")
	m.MacAddress = r.PostFormValue("macaddress")
	m.Profile = r.PostFormValue("profile")
	m.Save()
	http.Redirect(w, r, "/machines/"+m.Name, http.StatusMovedPermanently)
}

func DeleteMachineHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	DeleteMachineByName(vars["name"])
	http.Redirect(w, r, "/machines/", http.StatusMovedPermanently)
}

func CreateMachineHandler(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title: "Create Machine",
	}
	renderTemplate(w, "templates/machines/create.html", p)
}

func EditMachineHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	m, err := GetMachineByName(name)
	if err != nil {
		log.Println(err.Error())
	}
	p := &Page{
		Data:  m,
		Title: "Edit " + name,
	}
	renderTemplate(w, "templates/machines/edit.html", p)
}

func ShowMachineHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	m, err := GetMachineByName(name)
	if err != nil {
		log.Println(err.Error())
	}
	p := &Page{
		Data:  m,
		Title: m.Name,
	}
	renderTemplate(w, "templates/machines/machine.html", p)
}

func ListMachinesHandler(w http.ResponseWriter, r *http.Request) {
	machines, err := GetMachines()
	if err != nil {
		log.Println(err.Error())
	}
	p := &Page{
		Data:  machines,
		Title: "Machines",
	}
	renderTemplate(w, "templates/machines/list.html", p)
}
