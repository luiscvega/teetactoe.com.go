package logic

import (
	"log"

	"./../models"
)

func GetUser(id int64) (user *models.User, err error) {
        user = new(models.User)

        err = db.QueryRow("SELECT id, email, first_name, last_name FROM users WHERE id = $1", id).Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName)
        if err != nil {
                log.Fatal(err)
        }

        return
}
