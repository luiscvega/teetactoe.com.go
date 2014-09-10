package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"./logic"
	"./routes"
)

func main() {
	db, err := sql.Open("postgres", "postgres://localhost/luis?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	logic.DB = db

	r := mux.NewRouter()
	r.StrictSlash(true)

	routes.Initialize(r)

	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	http.Handle("/", r)

	http.ListenAndServe(":3000", nil)

	defer db.Close()
}
