package main

import (
	"log"
	"net/http"

	"github.com/kelseyhightower/ipxed/api"
	"github.com/kelseyhightower/ipxed/web"
)

func main() {
	apiHandler := api.Handler()
	webHandler := web.Handler()
	http.Handle("/api/", apiHandler)
	http.Handle("/", webHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
