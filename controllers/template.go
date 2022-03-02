package controllers

import "net/http"

// Template : template and data interface for controllers
type Template interface {
	Execute(w http.ResponseWriter, data interface{})
}
