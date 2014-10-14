package logic

import (
	"database/sql"
	"log"

	"./../models"
)

func GetUser(id int) (user *models.User, err error) {
	user = new(models.User)

	err = DB.QueryRow("SELECT id, email, first_name, last_name FROM users WHERE id = $1", id).Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}
