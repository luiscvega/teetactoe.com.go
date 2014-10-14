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

	page := locals.(Page)

	userId, ok := ctx.Session.Values["user_id"].(int)
	if ok {
		page.CurrentUser, _ = logic.GetUser(userId)
	}

	return t.Execute(ctx.Response, page)
}

func (ctx Context) Redirect(url string) {
	http.Redirect(ctx.Response, ctx.Request, url, 303)
}

type Page struct {
	CurrentUser  *models.User
	Locals       interface{}
	ErrorMessage string
}
