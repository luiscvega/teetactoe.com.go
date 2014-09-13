package routes

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/gorilla/sessions"

	"./../handlers"
	"./../logic"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func Initialize() *pat.PatternServeMux {
	m := pat.New()

	m.Get("/", prepare(handlers.RootGet))

	m.Get("/signup", prepare(handlers.SignupGet))
	m.Post("/signup", prepare(handlers.SignupPost))

	m.Get("/admin/logout", prepare(handlers.LogoutGet))
	m.Get("/admin/login", prepare(handlers.LoginGet))
	m.Post("/admin/login", prepare(handlers.LoginPost))

	m.Get("/admin/campaigns", prepare(handlers.CampaignsIndexGet))
	m.Post("/admin/campaigns", prepare(handlers.CampaignCreatePost))
	m.Get("/admin/campaigns/new", prepare(handlers.CampaignNewGet))

	return m
}

func prepare(handler func(ctx handlers.Context)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "teetactoe.com")

		page := handlers.Page{
			Session: session.Values}

		userId, ok := session.Values["user_id"].(int64)
		if ok {
			page.CurrentUser = logic.GetUser(userId)
		}

		ctx := handlers.Context{
			Response: w,
			Request:  r,
			Session:  session,
			Page:     page}

		handler(ctx)
	})
}
