package handlers

import (
	"./../forms"
	"./../logic"
)

func SignupGet(ctx Context) {
	ctx.Render("views/signup.html", ctx.Page)
}

func SignupPost(ctx Context) {
	password := ctx.Request.FormValue("password") // This calls r.ParseForm() already

	user, formErrors := forms.Signup.Validate(ctx.Request.Form)
	if formErrors.Any() {
		ctx.Render("views/signup.html", ctx.Page)
		return
	}

	err := logic.CreateUser(user, password)
	if err != nil {
		if err == logic.ErrUserExists {
			ctx.Render("views/signup.html", ctx.Page)
			return
		}

		panic(err)
	}

	ctx.Session.Values["user_id"] = user.Id
	ctx.Session.Save(ctx.Request, ctx.Response)

	ctx.Redirect("/")
}
