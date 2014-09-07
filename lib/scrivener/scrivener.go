package scrivener

import (
	"net/url"
)

type Form struct {
	Params url.Values
	Errors []string
}

func (f *Form) AssertPresent(field string) {
	if f.Params.Get(field) == "" {
		f.Errors = append(f.Errors, field+" can't be blank!")
	}
}

func (f *Form) AssertEqual(field1 string, field2 string) {
	if f.Params.Get(field1) != f.Params.Get(field2) {
		f.Errors = append(f.Errors, field1+" and "+field2+" do not match!")
	}
}
