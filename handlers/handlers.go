package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/sessions"

	"./../forms"
	"./../logic"
	"./../models"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func render(view string) *template.Template {
	return template.Must(template.ParseFiles("views/layout.html", view))
}

type Page struct {
	Session     map[interface{}]interface{}
	Flashes     []interface{}
	CurrentUser *models.User
}

func Root(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "teetactoe.com")
	flashes := session.Flashes()
	session.Save(r, w)

	user := new(models.User)
	userId, ok := session.Values["user_id"].(int64)
	if ok {
		user = logic.GetUser(userId)
	}

	page := Page{
		Session:     session.Values,
		Flashes:     flashes,
		CurrentUser: user}

	render("views/index.html").Execute(w, page)
}

func Signup(w http.ResponseWriter, r *http.Request) {
	render("views/signup.html").Execute(w, nil)
}

func SignupSubmit(w http.ResponseWriter, r *http.Request) {
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
}

func Login(w http.ResponseWriter, r *http.Request) {
	render("views/login.html").Execute(w, nil)
}

func LoginSubmit(w http.ResponseWriter, r *http.Request) {
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
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "teetactoe.com")
	delete(session.Values, "user_id")
	session.Save(r, w)

	http.Redirect(w, r, "/", 303)
}
