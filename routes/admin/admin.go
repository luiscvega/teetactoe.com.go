package admin

import (
	"log"
	"net/http"
	"text/template"

	"github.com/bmizerany/pat"
	"github.com/gorilla/sessions"

	"./../../forms"
	"./../../logic"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func Initialize() *pat.PatternServeMux {
        m := pat.New()

	m.Get("/admin/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("views/layout.html", "views/login.html"))
		t.Execute(w, nil)
	}))

	m.Post("/admin/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email := r.FormValue("email")
		password := r.FormValue("password")

		formErrors := forms.Login.Validate(r.Form)
		if len(formErrors) > 0 {
			log.Fatal(formErrors)
			return
		}

		user := logic.AuthenticateUser(email, password)

		session, _ := store.Get(r, "teetactoe.com")
		session.Values["user_id"] = user.Id
		session.Save(r, w)

		http.Redirect(w, r, "/", 303)
	}))

	m.Get("/admin/logout", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "teetactoe.com")
		delete(session.Values, "user_id")
		session.Save(r, w)

		http.Redirect(w, r, "/", 303)
	}))

        return m
}
