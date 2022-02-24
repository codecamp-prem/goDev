package controllers

import (
	"net/http"

	"github.com/codecamp-prem/goDev/views"
)

// StaticHandler : handle static template
func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
