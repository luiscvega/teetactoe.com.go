package routes

import (
	"fmt"
	"net/http"
	"text/template"

	"./../forms"
	"./../logic"
	"./admin"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

type Page struct {
	Session *sessions.Session
}

func Initialize(r *mux.Router) {
	r.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "teetactoe.com")
		t := template.Must(template.ParseFiles("views/layout.html", "views/index.html"))
		page := Page{session}
		fmt.Println(page.Session.Values)
		t.Execute(w, page)
	})).Methods("GET")

	r.HandleFunc("/signup", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("views/layout.html", "views/signup.html"))
		t.Execute(w, nil)
	})).Methods("GET")

	r.HandleFunc("/signup", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		password := r.FormValue("password") // This calls r.ParseForm() already

		user, formErrors := forms.Signup.Validate(r.Form)
		if formErrors.Any() {
			fmt.Println(formErrors)
			return
		}

		if err := logic.CreateUser(user, password); err != nil {
		}

		session, _ := store.Get(r, "teetactoe.com")
		session.Values["user_id"] = user.Id
		session.Save(r, w)

		http.Redirect(w, r, "/", 303)
	})).Methods("POST")

	admin.Initialize(r.PathPrefix("/admin").Subrouter())
}
