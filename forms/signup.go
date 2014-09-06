package forms

import (
	"./../models"

	"net/http"
	"net/url"
)

type Signup struct {
	Errors []string
	Params url.Values
}

func (s *Signup) Validate(r *http.Request) (err error, user *models.User) {
	r.ParseForm()

	user = new(models.User)

	s.Params = r.Form
	s.Errors = make([]string, 0)

	s.assertPresent("email")
	s.assertPresent("first_name")
	s.assertPresent("last_name")
	s.assertEqual("password", "confirm_password")

	return
}

func (s *Signup) assertPresent(field string) {
	if s.Params.Get(field) == "" {
		s.Errors = append(s.Errors, field + " can't be blank!")
	}
}

func (s *Signup) assertEqual(field1 string, field2 string) {
	if s.Params.Get(field1) != s.Params.Get(field2) {
		s.Errors = append(s.Errors, field1 + " and " + field2 + " do not match!")
	}
}
