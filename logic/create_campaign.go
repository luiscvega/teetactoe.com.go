package logic

import (
	"fmt"
	"log"

	"github.com/luiscvega/scrivener"

	"./../models"
)

func CreateCampaign(campaign *models.Campaign, userId int64) (err error) {
	campaign.UserId = userId

	scrivener := scrivener.New(campaign)
	scrivener.AssertPresent("UserId")

	if len(scrivener.Errors) > 0 {
		fmt.Println("HELLO!")
		log.Fatal("NO USER ID!")
		return
	}

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
