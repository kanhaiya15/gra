package kvalidator

import (
	"net/http"
	"net/url"

	"github.com/thedevsaddam/govalidator"
)

// Validate Validate returns nil or ValidationErrors
func Validate(r *http.Request, rules, messages map[string][]string) url.Values {
	opts := govalidator.Options{
		Request:  r,
		Rules:    rules,
		Messages: messages,
	}
	v := govalidator.New(opts)
	e := v.ValidateJSON()
	return e
}
