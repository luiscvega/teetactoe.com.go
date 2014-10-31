package main

import (
	"database/sql"
	"log"
	"net/http"

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

	http.Handle("/", routes.Guest())
	http.Handle("/admin/", http.StripPrefix("/admin", routes.Admin()))
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))

	http.ListenAndServe(":3000", nil)

	defer db.Close()
}
