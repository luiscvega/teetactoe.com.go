package routes

import (
	"html/template"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/bmizerany/pat"
	"github.com/gorilla/sessions"

	h "./../handlers"
)

type x struct {
	pat *pat.PatternServeMux
}

func (thing x) Get(path string, hand func(h.Context)) {
	thing.pat.Get(path, prepare(hand))
}

func (thing x) Post(path string, hand func(h.Context)) {
	thing.pat.Post(path, prepare(hand))
}

func (thing x) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	thing.pat.ServeHTTP(w,r)
}

func Admin() x {
	m := pat.New()

	y := x{m}

	y.Get("/logout", h.LogoutGet)
	y.Get("/login", h.LoginGet)
	y.Post("/login", h.LoginPost)

	y.Get("/campaigns", h.CampaignsIndexGet)
	y.Post("/campaigns", h.CampaignCreatePost)
	y.Get("/campaigns/new", h.CampaignNewGet)

	return y
}

func Guest() *pat.PatternServeMux {
	m := pat.New()

	m.Get("/", prepare(h.RootGet))

	m.Get("/signup", prepare(h.SignupGet))
	m.Post("/signup", prepare(h.SignupPost))

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

		defer catchPanic(w)

		handler(ctx)
	})
}

func catchPanic(w http.ResponseWriter) {
	if r := recover(); r != nil {
		log.Println("ERROR:", r)
		log.Printf("TRACE: %s", debug.Stack())
		t := template.Must(template.ParseFiles("views/404.html"))
		t.Execute(w, nil)
	}
}
