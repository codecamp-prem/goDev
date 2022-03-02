package controllers

import (
	"fmt"
	"net/http"
)

//Users to hold the method template ofr user
type Users struct {
	Templates struct {
		New Template
	}
}

// New : Take it to sign up form
func (u Users) New(w http.ResponseWriter, r *http.Request) {
	u.Templates.New.Execute(w, nil)
}

// Create : Parse the data from signup form
func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Email: ", r.FormValue("email"))
	fmt.Fprint(w, "Password: ", r.FormValue("password"))
}
