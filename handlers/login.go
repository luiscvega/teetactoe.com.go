package handlers

import (
	"log"

	"./../forms"
	"./../logic"
)

func LoginGet(ctx Context) {
	ctx.Render("views/login.html", ctx.Page)
}

func LoginPost(ctx Context) {
	email := ctx.Request.FormValue("email")
	password := ctx.Request.FormValue("password")

	formErrors := forms.Login.Validate(ctx.Request.Form)
	if len(formErrors) > 0 {
		ctx.Render("views/login.html", ctx.Page)
		return
	}

	user, err := logic.AuthenticateUser(email, password)
	if err != nil {
		log.Fatal(err)
	}

	if user == nil {
		// User was not found. Invalid credentials?
		ctx.Render("views/login.html", ctx.Page)
		return
	}

	ctx.Session.Values["user_id"] = user.Id
	ctx.Session.Save(ctx.Request, ctx.Response)

	ctx.Redirect("/")
}
