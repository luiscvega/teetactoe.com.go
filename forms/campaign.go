package forms

import (
	"net/url"

	"github.com/luiscvega/body"
	"github.com/luiscvega/scrivener"

	"./../models"
)

type campaign struct {
	Name string `name:"name"`
}

func (form *campaign) Validate(params url.Values) (campaign *models.Campaign, formErrors scrivener.Errors) {
	body.Parse(params, form)

	scrivener := scrivener.New(form)
	scrivener.AssertPresent("Name")

	if len(scrivener.Errors) > 0 {
		formErrors = scrivener.Errors
		return
	}

	campaign = new(models.Campaign)
	campaign.Name = form.Name

	return
}
