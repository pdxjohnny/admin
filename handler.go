package main

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

// MakeHandler creates the api request handler
func MakeHandler() *http.Handler {
	api := rest.NewApi()

	api.Use(rest.DefaultProdStack...)
	router, err := rest.MakeRouter(
		// To add a user
		rest.Get("/adduser/:username/:password", GetAddUser),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	handler := api.MakeHandler()
	return &handler
}
