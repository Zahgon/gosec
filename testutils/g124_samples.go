package testutils

import gosec "github.com/securego/gosec/v2"

var SampleCodeG124 = []CodeSample{

	{
		Code: []string{`
package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:  "session",
		Value: "abc123",
	}
	http.SetCookie(w, cookie)
}
`},
		Errors: 1,
		Config: gosec.NewConfig(),
	},

	{
		Code: []string{`
package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     "session",
		Value:    "abc123",
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, cookie)
}
`},
		Errors: 1,
		Config: gosec.NewConfig(),
	},

	{
		Code: []string{`
package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     "session",
		Value:    "abc123",
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, cookie)
}
`},
		Errors: 1,
		Config: gosec.NewConfig(),
	},

	{
		Code: []string{`
package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     "session",
		Value:    "abc123",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, cookie)
}
`},
		Errors: 0,
		Config: gosec.NewConfig(),
	},

	{
		Code: []string{`
package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     "session",
		Value:    "abc123",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, cookie)
}
`},
		Errors: 0,
		Config: gosec.NewConfig(),
	},
}
