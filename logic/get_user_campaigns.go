package logic

import (
	"log"

	"./../models"
)

func GetUserCampaigns(userId int64) (campaigns []*models.Campaign) {
	stmt, err := DB.Prepare("SELECT id, name FROM campaigns WHERE user_id = $1")
        if err != nil {
                log.Fatal(err)
        }

        rows, err := stmt.Query(userId)
	if err != nil {
		log.Fatal(err)
	}

        campaigns = make([]*models.Campaign, 0)
        for rows.Next() {
                campaign := models.Campaign{}
                rows.Scan(&campaign.Id, &campaign.Name)
                campaigns = append(campaigns, &campaign)
        }

	return
}
