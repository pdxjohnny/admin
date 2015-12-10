package main

import (
	"fmt"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

var blankresponse = []byte("{}")

// GetAddUser creates a users account from a get request
func GetAddUser(w rest.ResponseWriter, r *rest.Request) {
	username := r.PathParam("username")
	password := r.PathParam("password")
	fmt.Println("Adding user", username, password)
	w.WriteHeader(http.StatusOK)
	w.(http.ResponseWriter).Write(blankresponse)
}
