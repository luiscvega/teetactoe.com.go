package forms

import (
	"net/url"

	"github.com/luiscvega/body"
	"github.com/luiscvega/scrivener"

	"./../models"
)

type signup struct {
	Email           string `name:"email"`
	FirstName       string `name:"first_name"`
	LastName        string `name:"last_name"`
	Password        string `name:"password"`
	ConfirmPassword string `name:"confirm_password"`
}

func (form *signup) Validate(params url.Values) (user *models.User, formErrors scrivener.Errors) {
	body.Parse(params, form)

	scrivener := scrivener.New(form)
	scrivener.AssertPresent("Email")
	scrivener.AssertPresent("FirstName")
	scrivener.AssertPresent("LastName")

	if scrivener.AssertPresent("Password") && scrivener.AssertPresent("ConfirmPassword") {
		scrivener.Assert(func(interface{}) bool {
			return form.Password == form.ConfirmPassword
		}, "passwords", "not_match")
	}

	if len(scrivener.Errors) > 0 {
		formErrors = scrivener.Errors
		return
	}

	user = new(models.User)
	user.Email = form.Email
	user.FirstName = form.FirstName
	user.LastName = form.LastName
	user.CryptedPassword = "CdFH2da9dFKkPnu23782"

	return
}
