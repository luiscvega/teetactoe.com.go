package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"./routes"
)

func main() {
	r := mux.NewRouter()
	r.StrictSlash(true)

	routes.Register(r)

	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	http.Handle("/", r)

	http.ListenAndServe(":3000", nil)
}
