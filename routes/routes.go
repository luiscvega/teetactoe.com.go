package routes

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"./../forms"
	"./../logic"
	"./../models"

	"github.com/bmizerany/pat"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

type Page struct {
	Session     map[interface{}]interface{}
	Flashes     []interface{}
	CurrentUser *models.User
}

func Initialize() *pat.PatternServeMux {
	m := pat.New()

	m.Get("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	}))

	m.Get("/signup", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("views/layout.html", "views/signup.html"))
		t.Execute(w, nil)
	}))

	m.Post("/signup", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		password := r.FormValue("password") // This calls r.ParseForm() already

		user, formErrors := forms.Signup.Validate(r.Form)
		if formErrors.Any() {
			fmt.Println(formErrors)
			return
		}

		if err := logic.CreateUser(user, password); err != nil {
			switch {
			case err.Error() == "A user with that email already exists!":
				http.Error(w, err.Error(), 500)
				return
			default:
				log.Fatal(err)
			}
		}

		session, _ := store.Get(r, "teetactoe.com")
		session.Values["user_id"] = user.Id
		session.Save(r, w)

		http.Redirect(w, r, "/", 303)
	}))

	return m
}
