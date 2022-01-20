package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to Golang</h1>")
}
func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>Mail us at:<a href=\"mailto:prem.bikram.limbu@gmail.com\">prem.bikram.limbu@gmail.com</a>.")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<p>Q: Is there a free version?<br>A: Yes, we offer 30 days in any paid plans.</p><p>Q: What are your hours?<br>A: We support 24/7 hours, though response time may be slower during weekends.</p><p>Q: How do I contact support?<br>A: Email us : godev-care@godev.com</p>")
}

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		//http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 		http.Error(w, "Page Not Found!", http.StatusNotFound)
// 	}
// }

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		//http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		http.Error(w, "Page Not Found!", http.StatusNotFound)
	}
}
func main() {
	var router Router
	fmt.Println("Starting Server at 3000...")
	http.ListenAndServe(":3000", router)
}
