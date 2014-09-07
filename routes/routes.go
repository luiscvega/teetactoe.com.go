package routes

import (
	"fmt"
	"net/http"
	"text/template"

	"./../forms"
	"./../logic"
	"./admin"

	"github.com/gorilla/mux"
)

func Initialize(r *mux.Router) {
	r.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("views/layout.html", "views/index.html"))
		t.Execute(w, "index")
	})).Methods("GET")

	r.HandleFunc("/signup", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("views/layout.html", "views/signup.html"))
		t.Execute(w, nil)
	})).Methods("GET")

	r.HandleFunc("/signup", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		user, formErrors := forms.Signup.Validate(r.Form)
		if formErrors.Any() {
			fmt.Println(formErrors)
			return
		}

		if err := logic.CreateUser(user); err != nil {
		}
	})).Methods("POST")

	admin.Initialize(r.PathPrefix("/admin").Subrouter())
}
