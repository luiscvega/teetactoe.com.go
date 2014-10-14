package logic

import (
	"errors"
	"log"

	"./../models"
)

func CreateUser(user *models.User, password string) (err error) {
	user.CryptedPassword = password

	var count int

	err = DB.QueryRow("SELECT count(*) FROM users WHERE email = $1", user.Email).Scan(&count)
	if err != nil {
		log.Panic(err)
	}

	if count == 1 {
		return errors.New("A user with that email already exists!")
	}

	stmt, err := DB.Prepare("INSERT INTO users (email, first_name, last_name, crypted_password) VALUES ($1, $2, $3, $4) RETURNING id")
	if err != nil {
		log.Panic(err)
	}

	err = stmt.QueryRow(user.Email, user.FirstName, user.LastName, user.CryptedPassword).Scan(&user.Id)
	if err != nil {
		log.Panic(err)
	}

	return nil
}
