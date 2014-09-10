package admin

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

	"./../../forms"
	"./../../logic"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func Initialize(r *mux.Router) {
	r.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("views/layout.html", "views/login.html"))
		t.Execute(w, nil)
	})).Methods("GET")

	r.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email := r.FormValue("email")

		formErrors := forms.Login.Validate(r.Form)
		if len(formErrors) > 0 {
			log.Fatal(formErrors)
			return
		}

		user := logic.GetUserByEmail(email)

		session, _ := store.Get(r, "teetactoe.com")
		session.Values["user_id"] = user.Id
		session.Save(r, w)

		http.Redirect(w, r, "/", 303)
	})).Methods("POST")

	r.HandleFunc("/logout", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "teetactoe.com")
		delete(session.Values, "user_id")
		session.Save(r, w)

		http.Redirect(w, r, "/", 303)
	})).Methods("GET")
}
