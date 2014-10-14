package handlers

import (
	"fmt"

	"./../forms"
	"./../logic"
)

func CampaignsIndexGet(ctx Context) {
	campaigns := logic.GetUserCampaigns(ctx.Session.Values["user_id"].(int))
	ctx.Render("views/campaigns/index.html", map[string]interface{}{
		"Campaigns": campaigns})
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

	if err := logic.CreateCampaign(campaign, ctx.Session.Values["user_id"].(int)); err != nil {
	}

	ctx.Redirect("/admin/campaigns")
}
