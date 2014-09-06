package routes

import (
	"errors"
	"net/http"
	"text/template"

	"./../logic"
	"./../models"
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
		err, user := validateSignupForm(r)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		if err = logic.CreateUser(user); err != nil {
		}
	})).Methods("POST")

	admin.Routes(r.PathPrefix("/admin").Subrouter())
}

func validateSignupForm(r *http.Request) (err error, user *models.User) {
	r.ParseForm()

	user = new(models.User)
	form := r.Form

	if form.Get("password") != form.Get("confirm_password") {
		err = errors.New("Passwords don't match!!!")
		return
	}

	if user.Email = form.Get("email"); user.Email == "" {
		err = errors.New("Email can't be blank!")
		return
	}

	if user.FirstName = form.Get("first_name"); user.FirstName == "" {
		err = errors.New("First name can't be blank!")
		return
	}

	if user.LastName = form.Get("last_name"); user.LastName == "" {
		err = errors.New("Last name can't be blank!")
		return
	}

	return
}
