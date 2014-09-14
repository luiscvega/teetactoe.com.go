package handlers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"

	"./../logic"
	"./../models"
)

type Context struct {
	Response http.ResponseWriter
	*http.Request
	*sessions.Session
	Page
}

func (ctx Context) Render(view string, locals interface{}) {
	t := template.Must(template.ParseFiles("views/layout.html", view))

	page := Page{
		Locals: locals}

	userId, ok := ctx.Session.Values["user_id"].(int64)
	if ok {
		page.CurrentUser = logic.GetUser(userId)
	}

	t.Execute(ctx.Response, page)
}

func (ctx Context) Redirect(url string) {
	http.Redirect(ctx.Response, ctx.Request, url, 303)
}

type Page struct {
	CurrentUser *models.User
	Locals      interface{}
}
