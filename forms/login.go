package forms

import (
	"./../lib/scrivener"

	"net/url"
)

func ValidateLoginForm(params url.Values) []string {
	f := new(scrivener.Form)
	f.Errors = make([]string, 0)
	f.Params = params

	f.AssertPresent("email")
	f.AssertPresent("password")

	return f.Errors
}
