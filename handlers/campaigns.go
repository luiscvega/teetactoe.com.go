package handlers

import (

	"./../forms"
	"./../logic"
)

func CampaignsIndexGet(ctx Context) {
	campaigns, _ := logic.GetUserCampaigns(ctx.Session.Values["user_id"].(int))
	ctx.Render("views/campaigns/index.html", map[string]interface{}{
		"Campaigns": campaigns})
}

func CampaignNewGet(ctx Context) {
	ctx.Render("views/campaigns/new.html", ctx.Page)
}

func CampaignsShowGet(ctx Context) {
	campaignId := ctx.Request.URL.Query().Get(":campaign_id")
	campaign := logic.GetCampaign(campaignId)
	ctx.Render("views/campaigns/show.html", map[string]interface{}{
		"Campaign": campaign})
}

func CampaignCreatePost(ctx Context) {
	ctx.ParseForm()

	campaign, formErrors := forms.Campaign.Validate(ctx.Request.Form)
	if formErrors.Any() {
		return
	}

	if err := logic.CreateCampaign(campaign, ctx.Session.Values["user_id"].(int)); err != nil {
	}

	ctx.Redirect("/admin/campaigns")
}
