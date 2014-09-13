package handlers

import (
	"fmt"

	"./../forms"
	"./../logic"
)

func CampaignsIndexGet(ctx Context) {
	ctx.Render("views/campaigns/index.html", ctx.Page)
}

func CampaignNewGet(ctx Context) {
	ctx.Render("views/campaigns/new.html", ctx.Page)
}

func CampaignCreatePost(ctx Context) {
	ctx.ParseForm()

	campaign, formErrors := forms.CreateCampaign.Validate(ctx.Request.Form)
	if formErrors.Any() {
		fmt.Println(formErrors)
		return
	}

	if err := logic.CreateCampaign(campaign, ctx.Page.CurrentUser.Id); err != nil {
	}

	ctx.Redirect("/")
}
