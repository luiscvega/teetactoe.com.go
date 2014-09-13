package handlers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"

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
	t.Execute(ctx.Response, locals)
}

func (ctx Context) Redirect(url string) {
	http.Redirect(ctx.Response, ctx.Request, url, 303)
}

type Page struct {
	Session     map[interface{}]interface{}
	CurrentUser *models.User
}
