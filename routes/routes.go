package routes

import (
	"net/http"

	"github.com/bmizerany/pat"

	"./../handlers"
)

func Initialize() *pat.PatternServeMux {
	m := pat.New()

	m.Get("/", http.HandlerFunc(handlers.Root))
	m.Get("/signup", http.HandlerFunc(handlers.Signup))
	m.Post("/signup", http.HandlerFunc(handlers.SignupSubmit))

	m.Get("/admin/", http.HandlerFunc(handlers.Login))
	m.Post("/admin/", http.HandlerFunc(handlers.LoginSubmit))
	m.Get("/admin/logout", http.HandlerFunc(handlers.Logout))

	return m
}
