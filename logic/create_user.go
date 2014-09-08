package logic

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"./../models"
)

func CreateUser(user *models.User, password string) (err error) {
	db, err := sql.Open("postgres", "postgres://localhost/luis?sslmode=disable")

	user.CryptedPassword = "booomboxhawehfajkdfhaljkshr3lhalf"

	stmt, err := db.Prepare("INSERT INTO users (email, first_name, last_name, password) VALUES ($1, $2, $3, $4) RETURNING id")
	if err != nil {
		log.Fatal(err)
	}

	err = stmt.QueryRow(user.Email, user.FirstName, user.LastName, user.CryptedPassword).Scan(&user.Id)
	if err != nil {
		log.Fatal(err)
		return err
	}

	defer db.Close()

	return
}
