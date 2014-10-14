package logic

import (
	"database/sql"
	"errors"
	"log"

	"./../models"
)

func AuthenticateUser(email, password string) (user *models.User, err error) {
	user = new(models.User)

	err = DB.QueryRow("SELECT id, email, first_name, last_name FROM users WHERE email = $1 AND crypted_password = $2", email, password).Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName)
	switch {
	case err == sql.ErrNoRows:
		err = errors.New("User with those credentials not found!")
	case err != nil:
		log.Fatal(err)
	}

	return
}
