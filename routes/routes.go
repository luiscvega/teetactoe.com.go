package routes

import (
	"fmt"
	"net/http"
	"text/template"

	"./../logic"
	"./../forms"
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
		signupForm := new(forms.Signup)
		err, user := signupForm.Validate(r)

		if len(signupForm.Errors) > 0 {
			fmt.Println(signupForm.Errors)
		}

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		if err = logic.CreateUser(user); err != nil {
		}
	})).Methods("POST")

	admin.Routes(r.PathPrefix("/admin").Subrouter())
}
