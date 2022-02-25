package controllers

import (
	"html/template"
	"net/http"

	"github.com/codecamp-prem/goDev/views"
)

// StaticHandler : handle static template
func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

// FAQ : dynamic faq
func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "Is there a free version?",
			Answer:   "Yes, we offer 30 days in any paid plans.",
		},
		{
			Question: "What are your hours?",
			Answer:   "We support 24/7 hours, though response time may be slower during weekends.",
		},
		{
			Question: "How do I contact support?",
			Answer:   `Email us : <a href="mailto:support@google.com">godev_care@godev.com</a>`,
		},
		{
			Question: "Any hidden charges?",
			Answer:   "No! It is what you see in pricing list.",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
