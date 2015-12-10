package main

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/pdxjohnny/numapp/api/recaptcha"
	"github.com/spf13/viper"
)

var blankresponse = []byte("{}")

// GetAddUser creates a users account from a get request
func GetAddUser(w rest.ResponseWriter, r *rest.Request) {
	username := r.PathParam("username")
	password := r.PathParam("password")
	reCAPTCHA := r.PathParam("recaptcha")
	fmt.Println("Adding user", username, password)
	// Verify with google reCAPTCHA
	err := recaptcha.Verify(viper.GetString("recaptcha"), reCAPTCHA)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = AddUser(username, password)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.(http.ResponseWriter).Write(blankresponse)
}

// AddUser creates a users account
func AddUser(username, password string) error {
	cmd := exec.Command("adduser", username)
	err := cmd.Start()
	if err != nil {
		return err
	}
	err = cmd.Wait()
	if err != nil {
		return err
	}
	userpass := fmt.Sprintf("%s:%s", username, password)
	cmd = exec.Command("echo", userpass, "|", "chpasswd")
	err = cmd.Start()
	if err != nil {
		return err
	}
	err = cmd.Wait()
	if err != nil {
		return err
	}
	return nil
}
