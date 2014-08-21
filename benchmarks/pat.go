package main

import (
	"net/http"
	"text/template"

	"github.com/bmizerany/pat"
)

func main() {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("views/layout.html", "views/index.html"))
		t.Execute(w, nil)
	}))

	mux.Get("/signup", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("views/layout.html", "views/signup.html"))
		t.Execute(w, nil)
	}))

	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	http.Handle("/", mux)

	http.ListenAndServe(":3000", nil)
}
