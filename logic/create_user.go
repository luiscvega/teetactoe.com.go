package logic

import (
	"log"

	"./../models"
)

func CreateUser(user *models.User, password string) (err error) {
	user.CryptedPassword = "booomboxhawehfajkdfhaljkshr3lhalf"

	stmt, err := DB.Prepare("INSERT INTO users (email, first_name, last_name, crypted_password) VALUES ($1, $2, $3, $4) RETURNING id")
	if err != nil {
		log.Fatal(err)
	}

	err = stmt.QueryRow(user.Email, user.FirstName, user.LastName, user.CryptedPassword).Scan(&user.Id)
	if err != nil {
		log.Fatal(err)
	}

	return
}
