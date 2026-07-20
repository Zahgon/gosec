package testutils

import "github.com/securego/gosec/v2"

var SampleCodeG121 = []CodeSample{

	{[]string{`
package main

import "net/http"

func setup() {
	var cop http.CrossOriginProtection
	cop.AddInsecureBypassPattern("/")
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "net/http"

func setup() {
	var cop http.CrossOriginProtection
	cop.AddInsecureBypassPattern("/*")
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	_ = w
	var cop http.CrossOriginProtection
	pattern := r.URL.Query().Get("bypass")
	cop.AddInsecureBypassPattern(pattern)
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "net/http"

func setup() {
	var cop http.CrossOriginProtection
	cop.AddInsecureBypassPattern("/healthz")
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "net/http"

func setup() {
	var cop http.CrossOriginProtection
	cop.AddInsecureBypassPattern("/status")
	cop.AddInsecureBypassPattern("/metrics")
}
`}, 0, gosec.NewConfig()},
}
