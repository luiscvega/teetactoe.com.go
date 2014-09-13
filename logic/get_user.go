package logic

import (
	"database/sql"
	"log"

	"./../models"
)

func GetUser(id int64) (user *models.User) {
	user = new(models.User)

	err := DB.QueryRow("SELECT id, email, first_name, last_name FROM users WHERE id = $1", id).Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName)
	switch {
	case err == sql.ErrNoRows:
		log.Println("No rows!")
		err = nil
	case err != nil:
		log.Fatal(err)
	}

	user.Campaigns = GetUserCampaigns(user.Id)

	return
}
