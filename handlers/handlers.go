package handlers

import (
	"fmt"
	"log"
	"net/http"

	"./../forms"
	"./../logic"
)

func Root(ctx Context) {
	ctx.Render("views/index.html", ctx.Page)
}

func Signup(ctx Context) {
	ctx.Render("views/signup.html", ctx.Page)
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

	ctx.Redirect("/")
}

func Login(ctx Context) {
	ctx.Render("views/login.html", ctx.Page)
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

	ctx.Redirect("/")
}

func Logout(ctx Context) {
	delete(ctx.Session.Values, "user_id")
	ctx.Session.Save(ctx.Request, ctx.Response)

	ctx.Redirect("/")
}
