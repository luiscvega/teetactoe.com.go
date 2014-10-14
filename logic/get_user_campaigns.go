package logic

import (
	"./../models"
)

func GetUserCampaigns(userId int) (campaigns []*models.Campaign, err error) {
	stmt, err := DB.Prepare("SELECT id, name FROM campaigns WHERE user_id = $1")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(userId)
	if err != nil {
		return nil, err
	}

	campaigns = make([]*models.Campaign, 0)
	for rows.Next() {
		campaign := models.Campaign{}
		rows.Scan(&campaign.Id, &campaign.Name)
		campaigns = append(campaigns, &campaign)
	}

	return campaigns, nil
}
