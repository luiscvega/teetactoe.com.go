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

func Routes(r *mux.Router) {
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

		user, formErrors := forms.ValidateSignupForm(r.Form)
		if len(formErrors) > 0 {
			fmt.Println(formErrors)
			return
		}

		if err := logic.CreateUser(user); err != nil {
		}
	})).Methods("POST")

	r.HandleFunc("/login", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("views/layout.html", "views/login.html"))
		t.Execute(w, nil)
	})).Methods("GET")

	r.HandleFunc("/login", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		formErrors := forms.ValidateLoginForm(r.Form)
		if len(formErrors) > 0 {
			fmt.Println(formErrors)
			return
		}

		fmt.Println("No errors!")
	})).Methods("POST")

	admin.Routes(r.PathPrefix("/admin").Subrouter())
}
