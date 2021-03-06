package logic

import (
	"database/sql"
	"log"

	"./../models"
)

func GetUser(userId interface{}) (user *models.User) {
	user = new(models.User)

	err := DB.QueryRow("SELECT id, email, first_name, last_name FROM users WHERE id = $1", userId).Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}

		log.Panic(err)
	}

	return user
}
