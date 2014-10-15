package handlers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"

	"./../logic"
	"./../models"
)

type Context struct {
	Locals   map[string]interface{}
	Response http.ResponseWriter
	*http.Request
	*sessions.Session
	Page
}

func (ctx Context) Render(view string, locals interface{}) error {
	t := template.Must(template.ParseFiles("views/layout.html", view))

	ctx.Page.Locals = locals
	ctx.Page.CurrentUser = logic.GetUser(ctx.Session.Values["user_id"])

	return t.Execute(ctx.Response, ctx.Page)
}

func (ctx Context) Redirect(url string) {
	http.Redirect(ctx.Response, ctx.Request, url, 303)
}

type Page struct {
	CurrentUser  *models.User
	Locals       interface{}
	ErrorMessage string
}
