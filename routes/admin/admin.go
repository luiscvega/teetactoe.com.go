package admin

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"

	"./../../forms"
)

func Initialize(r *mux.Router) {
	r.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("views/layout.html", "views/login.html"))
		t.Execute(w, nil)
	})).Methods("GET")

	r.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		formErrors := forms.Login.Validate(r.Form)
		if len(formErrors) > 0 {
			fmt.Println(formErrors)
			return
		}
		fmt.Println("No errors!")
	})).Methods("POST")
}
