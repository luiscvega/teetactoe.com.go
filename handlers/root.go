package handlers

func RootGet(ctx Context) {
	ctx.Render("views/index.html", ctx.Page)
}
