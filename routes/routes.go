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

	m.Get("/", prepare(handlers.Root))

	m.Get("/signup", prepare(handlers.Signup))
	m.Post("/signup", prepare(handlers.SignupSubmit))

	m.Get("/admin/logout", prepare(handlers.Logout))
	m.Get("/admin/", prepare(handlers.Login))
	m.Post("/admin/", prepare(handlers.LoginSubmit))

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
