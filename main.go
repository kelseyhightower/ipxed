package main

import (
	"fmt"
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
	port := "0.0.0.0:8080"
	fmt.Println("Listening on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
