package forms

import (
	"./../lib/scrivener"
	"./../models"

	"net/url"
)

func ValidateSignupForm(params url.Values) (user *models.User, formErrors []string) {
	f := new(scrivener.Form)
	f.Errors = make([]string, 0)
	f.Params = params

	if !assertions(f) {
		formErrors = f.Errors
		return
	}

	user = run(f)

	return
}

func run(f *scrivener.Form) *models.User {
	params := f.Params

	user := new(models.User)
	user.Email = params.Get("email")
	user.LastName = params.Get("last_name")
	user.FirstName = params.Get("first_name")
	user.CryptedPassword = "CdFH2da9dFKkPnu23782"

	return user
}

func assertions(f *scrivener.Form) bool {
	f.AssertPresent("email")
	f.AssertPresent("first_name")
	f.AssertPresent("last_name")
	f.AssertEqual("password", "confirm_password")

	if len(f.Errors) > 0 {
		return false
	}

	return true
}
