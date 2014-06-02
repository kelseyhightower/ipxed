package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/httputils"
)

func CreateProfileHandler(w http.ResponseWriter, r *http.Request) {
	var p Profile
	if err := httputils.UnmarshalJSONBody(r.Body, &p); err != nil {
		httputils.JSONError(w, err.Error(), 500)
	}
	if err := p.Save(); err != nil {
		httputils.JSONError(w, err.Error(), 500)
	}
	httputils.SetLocationHeader(w, "http://localhost:8080/profiles/"+p.Name)
	w.WriteHeader(http.StatusCreated)
}

func GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	p, err := GetProfileByName(vars["name"])
	if err == ErrNotFound {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		httputils.JSONError(w, err.Error(), 500)
		return
	}
	httputils.JSONWrite(w, p, 200)
}

type ProfilesResponse struct {
	Profiles []Profile `json:"profiles"`
}

func ListProfilesHandler(w http.ResponseWriter, r *http.Request) {
	var response ProfilesResponse
	profiles, err := GetProfiles()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
	}
	response.Profiles = profiles
	httputils.JSONWrite(w, response, 200)
}

func DeleteProfileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := DeleteProfileByName(vars["name"])
	if err != nil {
		httputils.JSONError(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
