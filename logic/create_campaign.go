package logic

import (
	"log"

	"./../models"
)

func CreateCampaign(campaign *models.Campaign, userId int64) (err error) {
	campaign.UserId = userId

	stmt, err := DB.Prepare("INSERT INTO campaigns (name, user_id) VALUES ($1, $2) RETURNING id")
	if err != nil {
		log.Fatal(err)
	}

	err = stmt.QueryRow(campaign.Name, campaign.UserId).Scan(&campaign.Id)
	if err != nil {
		log.Fatal(err)
	}

	return
}
