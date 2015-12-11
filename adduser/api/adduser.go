package api

import (
	"errors"
	"fmt"
	"log"
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
	if viper.GetString("recaptcha") != "" {
		fmt.Println("Verifiying reCAPTCHA for", username)
		reCAPTCHA := r.PathParam("recaptcha")
		// Verify with google reCAPTCHA
		err := recaptcha.Verify(viper.GetString("recaptcha"), reCAPTCHA)
		if err != nil {
			log.Println(err)
			rest.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	fmt.Println("Adding user", username)
	err := AddUser(username, password)
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.(http.ResponseWriter).Write(blankresponse)
}

// AddUser creates a users account
func AddUser(username, password string) error {
	out, err := exec.Command("useradd", "-m", username).CombinedOutput()
	if err != nil {
		return errors.New(string(out))
	}
	userpass := fmt.Sprintf("echo %s:%s | chpasswd", username, password)
	out, err = exec.Command("bash", "-c", userpass).CombinedOutput()
	if err != nil {
		return errors.New(string(out))
	}
	return nil
}
