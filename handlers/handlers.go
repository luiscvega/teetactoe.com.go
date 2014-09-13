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

func render(view string) *template.Template {
	return template.Must(template.ParseFiles("views/layout.html", view))
}

type Page struct {
	Session     map[interface{}]interface{}
	CurrentUser *models.User
}

type Context struct {
	Response http.ResponseWriter
	Request  *http.Request
	Session  *sessions.Session
}

func Root(ctx Context) {
	user := new(models.User)
	userId, ok := ctx.Session.Values["user_id"].(int64)
	if ok {
		user = logic.GetUser(userId)
	}

	page := Page{
		Session:     ctx.Session.Values,
		CurrentUser: user}

	render("views/index.html").Execute(ctx.Response, page)
}

func Signup(ctx Context) {
	render("views/signup.html").Execute(ctx.Response, nil)
}

func SignupSubmit(ctx Context) {
	password := ctx.Request.FormValue("password") // This calls r.ParseForm() already

	user, formErrors := forms.Signup.Validate(ctx.Request.Form)
	if formErrors.Any() {
		fmt.Println(formErrors)
		return
	}

	if err := logic.CreateUser(user, password); err != nil {
		switch {
		case err.Error() == "A user with that email already exists!":
			http.Error(ctx.Response, err.Error(), 500)
			return
		default:
			log.Fatal(err)
		}
	}

	ctx.Session.Values["user_id"] = user.Id
	ctx.Session.Save(ctx.Request, ctx.Response)

	http.Redirect(ctx.Response, ctx.Request, "/", 303)
}

func Login(ctx Context) {
	render("views/login.html").Execute(ctx.Response, nil)
}

func LoginSubmit(ctx Context) {
	email := ctx.Request.FormValue("email")
	password := ctx.Request.FormValue("password")

	formErrors := forms.Login.Validate(ctx.Request.Form)
	if len(formErrors) > 0 {
		log.Fatal(formErrors)
		return
	}

	user := logic.AuthenticateUser(email, password)

	ctx.Session.Values["user_id"] = user.Id
	ctx.Session.Save(ctx.Request, ctx.Response)

	http.Redirect(ctx.Response, ctx.Request, "/", 303)
}

func Logout(ctx Context) {
	delete(ctx.Session.Values, "user_id")
	ctx.Session.Save(ctx.Request, ctx.Response)

	http.Redirect(ctx.Response, ctx.Request, "/", 303)
}
