package forms

import (
	"net/url"

	"github.com/luiscvega/body"

	"./../lib/scrivener"
)

type Login struct {
	Email    string `name:"email"`
	Password string `name:"password"`
}

func (login *Login) Validate(params url.Values) []string {
	body.Parse(params, login)
	scrivener := scrivener.New(login)
	scrivener.AssertPresent("Email")
	scrivener.AssertPresent("Password")
	return scrivener.Errors
}
