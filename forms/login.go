package forms

import (
	"net/url"

	"github.com/luiscvega/body"

	"./../lib/scrivener"
)

type login struct {
	Email    string `name:"email"`
	Password string `name:"password"`
}

func (form *login) Validate(params url.Values) []string {
	body.Parse(params, form)
	scrivener := scrivener.New(form)
	scrivener.AssertPresent("Email")
	scrivener.AssertPresent("Password")
	return scrivener.Errors
}
