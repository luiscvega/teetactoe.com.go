package handlers

import (
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
	if formErrors.Any() {
		ctx.Render("views/login.html", ctx.Page)
		return
	}

	user := logic.AuthenticateUser(email, password)
	if user == nil {
		ctx.Page.ErrorMessage = "Invalid credentials!"
		ctx.Render("views/login.html", ctx.Page)
		return
	}

	ctx.Session.Values["user_id"] = user.Id
	ctx.Session.Save(ctx.Request, ctx.Response)

	ctx.Redirect("/")
}
