package routes

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"./../forms"
	"./../logic"
	"./../models"
	"./admin"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

type Page struct {
	Session     map[interface{}]interface{}
	Flashes     []interface{}
	CurrentUser *models.User
}

func Initialize(r *mux.Router) {
	r.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "teetactoe.com")
		flashes := session.Flashes()
		session.Save(r, w)

		t := template.Must(template.ParseFiles("views/layout.html", "views/index.html"))

		user := new(models.User)
		userId, ok := session.Values["user_id"].(int64)
		if ok {
			user = logic.GetUser(userId)
		}

		page := Page{
			Session:     session.Values,
			Flashes:     flashes,
			CurrentUser: user}

		t.Execute(w, page)
	})).Methods("GET")

	r.HandleFunc("/signup", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("views/layout.html", "views/signup.html"))
		t.Execute(w, nil)
	})).Methods("GET")

	r.HandleFunc("/signup", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "teetactoe.com")

		password := r.FormValue("password") // This calls r.ParseForm() already

		user, formErrors := forms.Signup.Validate(r.Form)
		if formErrors.Any() {
			fmt.Println(formErrors)
			return
		}

		if err := logic.CreateUser(user, password); err != nil {
			switch {
			case err.Error() == "A user with that email already exists!":
				session.AddFlash(err.Error())
				session.Save(r, w)
			default:
				log.Fatal(err)
			}
		}

		http.Redirect(w, r, "/", 303)
	})).Methods("POST")

	admin.Initialize(r.PathPrefix("/admin").Subrouter())
}
