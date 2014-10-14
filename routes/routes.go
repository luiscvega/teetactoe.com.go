package routes

import (
	"html/template"
	"log"
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/gorilla/sessions"

	h "./../handlers"
)

func Initialize() *pat.PatternServeMux {
	m := pat.New()

	m.Get("/", prepare(h.RootGet))

	m.Get("/signup", prepare(h.SignupGet))
	m.Post("/signup", prepare(h.SignupPost))

	m.Get("/admin/logout", prepare(h.LogoutGet))
	m.Get("/admin/login", prepare(h.LoginGet))
	m.Post("/admin/login", prepare(h.LoginPost))

	m.Get("/admin/campaigns", prepare(h.CampaignsIndexGet))
	m.Post("/admin/campaigns", prepare(h.CampaignCreatePost))
	m.Get("/admin/campaigns/new", prepare(h.CampaignNewGet))

	return m
}

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func prepare(handler func(ctx h.Context)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "teetactoe.com")

		ctx := h.Context{
			Response: w,
			Request:  r,
			Session:  session}

		defer catchPanic()

		handler(ctx)
	})
}

func catchPanic() {
	if r := recover(); r != nil {
		log.Println("ERROR:", r)
		t := template.Must(template.ParseFiles("views/404.html"))
		t.Execute(w, nil)
	}
}
