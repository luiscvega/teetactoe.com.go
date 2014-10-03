package handlers

func LogoutGet(ctx Context) {
	delete(ctx.Session.Values, "user_id")
	ctx.Session.Save(ctx.Request, ctx.Response)

	ctx.Redirect("/")
}
