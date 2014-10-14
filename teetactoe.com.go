package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"./logic"
	"./routes"
)

func init() {
	// Initialize DB
	if db, err := sql.Open("postgres", "postgres://localhost/luis?sslmode=disable"); err != nil {
		log.Fatal(err)
	}
	logic.DB = db
}

func main() {
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	http.Handle("/", routes.Initialize())
	http.ListenAndServe(":3000", nil)
	defer db.Close()
}
