package logic

import (
	"database/sql"
	"log"

	"./../models"
)

func GetCampaign(campaignId interface{}) *models.Campaign {
	campaign := new(models.Campaign)

	err := DB.QueryRow("SELECT id, name FROM campaigns WHERE id = $1", campaignId).Scan(&campaign.Id, &campaign.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}

		log.Panic(err)
	}

	return campaign
}
