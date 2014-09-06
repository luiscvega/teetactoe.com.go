package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://localhost/luis?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	var (
		first_name string
		last_name  string
	)

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&first_name, &last_name); err != nil {
			log.Fatal(err)
		}

		log.Println(first_name, last_name)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
