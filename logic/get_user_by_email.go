package logic

import (
	"database/sql"
	"log"

	"./../models"
)

func GetUserByEmail(email string) (user *models.User) {
	user = new(models.User)

	err := DB.QueryRow("SELECT id, email, first_name, last_name FROM users WHERE email = $1", email).Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName)
	switch {
	case err == sql.ErrNoRows:
		log.Println("No rows!")
		err = nil
	case err != nil:
		log.Fatal(err)
	}

	return
}
