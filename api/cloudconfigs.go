package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/httputils"
)

func CreateCloudConfigHandler(w http.ResponseWriter, r *http.Request) {
	var c CloudConfig
	if err := httputils.UnmarshalJSONBody(r.Body, &c); err != nil {
		httputils.JSONError(w, err.Error(), 500)
	}
	if err := c.Save(); err != nil {
		httputils.JSONError(w, err.Error(), 500)
	}
	httputils.SetLocationHeader(w, "http://localhost:8080/cloudconfigs/"+c.Name)
	w.WriteHeader(http.StatusCreated)
}

func GetCloudConfigHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	c, err := GetCloudConfigByName(vars["name"])
	if err == ErrNotFound {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		httputils.JSONError(w, err.Error(), 500)
		return
	}
	httputils.JSONWrite(w, c, 200)
}

type CloudConfigsResponse struct {
	CloudConfigs []CloudConfig `json:"cloudconfigs"`
}

func ListCloudConfigHandler(w http.ResponseWriter, r *http.Request) {
	var response CloudConfigsResponse
	cloudconfigs, err := GetCloudConfigs()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
	}
	response.CloudConfigs = cloudconfigs
	httputils.JSONWrite(w, response, 200)
}

func DeleteCloudConfigHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := DeleteCloudConfigByName(vars["name"])
	if err != nil {
		httputils.JSONError(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
