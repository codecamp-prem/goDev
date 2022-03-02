package controllers

import (
	"net/http"
)

//Users to hold the method template ofr user
type Users struct {
	Templates struct {
		New Template
	}
}

// New crete new user
func (u Users) New(w http.ResponseWriter, r *http.Request) {
	u.Templates.New.Execute(w, nil)
}
