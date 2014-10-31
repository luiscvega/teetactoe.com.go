package routes

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/gorilla/sessions"

	h "./../handlers"
)

func Admin() *pat.PatternServeMux {
	m := pat.New()

	m.Get("/logout", prepare(h.LogoutGet))
	m.Get("/login", prepare(h.LoginGet))
	m.Post("/login", prepare(h.LoginPost))

	m.Get("/campaigns", prepare(h.CampaignsIndexGet))
	m.Post("/campaigns", prepare(h.CampaignCreatePost))
	m.Get("/campaigns/new", prepare(h.CampaignNewGet))

	return m
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

		handler(ctx)
	})
}
