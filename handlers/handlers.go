package handlers

import (
	"fmt"
	"log"
	"net/http"

	"./../forms"
	"./../logic"
	"./../models"
)

type Page struct {
	Session     map[interface{}]interface{}
	CurrentUser *models.User
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

	ctx.render("views/index.html", page)
}

func Signup(ctx Context) {
	ctx.render("views/signup.html", nil)
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

	ctx.redirect("/")
}

func Login(ctx Context) {
	ctx.render("views/login.html", nil)
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

	ctx.redirect("/")
}

func Logout(ctx Context) {
	delete(ctx.Session.Values, "user_id")
	ctx.Session.Save(ctx.Request, ctx.Response)


	ctx.redirect("/")
}
