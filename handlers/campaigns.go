package handlers

import (
	"fmt"

	"./../forms"
	"./../logic"
	"./../models"
)

type RootPage struct {
	Campaigns []*models.Campaign
}

func CampaignsIndexGet(ctx Context) {
	campaigns := logic.GetUserCampaigns(ctx.Session.Values["user_id"].(int64))

	rootPage := RootPage{
		Campaigns: campaigns}

	ctx.Render("views/campaigns/index.html", rootPage)
}

func CampaignNewGet(ctx Context) {
	ctx.Render("views/campaigns/new.html", ctx.Page)
}

func CampaignCreatePost(ctx Context) {
	ctx.ParseForm()

	campaign, formErrors := forms.Campaign.Validate(ctx.Request.Form)
	if formErrors.Any() {
		fmt.Println(formErrors)
		return
	}

	if err := logic.CreateCampaign(campaign, ctx.Session.Values["user_id"].(int64)); err != nil {
	}

	ctx.Redirect("/")
}
