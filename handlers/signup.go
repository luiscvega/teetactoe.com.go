package handlers

import (
	"fmt"
	"log"
	"net/http"

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
