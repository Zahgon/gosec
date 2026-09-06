package testutils

import "github.com/securego/gosec/v2"

var SampleCodeG710 = []CodeSample{

	{[]string{`
package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("next")
	http.Redirect(w, r, target, http.StatusFound)
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dest := r.FormValue("redirect")
	http.Redirect(w, r, "/proxy?to="+dest, http.StatusSeeOther)
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/dashboard", http.StatusFound)
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "bad id", http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/users/%d", id), http.StatusFound)
}
`}, 0, gosec.NewConfig()},
}
