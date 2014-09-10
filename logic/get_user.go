package logic

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"./../models"
)

func GetUser(id int64) (user *models.User, err error) {
	db, err := sql.Open("postgres", "postgres://localhost/luis?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

        user = new(models.User)

        err = db.QueryRow(
                "SELECT id, email, first_name, last_name FROM users WHERE id = $1", id).Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName)
        if err != nil {
                log.Fatal(err)
        }

        return
}
