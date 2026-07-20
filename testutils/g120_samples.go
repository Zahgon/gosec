package testutils

import "github.com/securego/gosec/v2"

var SampleCodeG120 = []CodeSample{

	{[]string{`
package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	_ = w
	_ = r.ParseMultipartForm(32 << 20)
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	_ = w
	_ = r.ParseForm()
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	_ = w
	_ = r.FormValue("q")
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	_ = w
	_ = r.PostFormValue("q")
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	_ = r.ParseMultipartForm(32 << 20)
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	_ = w
	processUpload(r)
}

func processUpload(r *http.Request) {
	_ = r.ParseMultipartForm(32 << 20)
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "net/http"

func fooHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = formParser(r)
	_, _ = w.Write([]byte("foo"))
}

func formParser(r *http.Request) (string, error) {
	if err := r.ParseForm(); err != nil {
		return "", err
	}
	return r.FormValue("varName"), nil
}
`}, 0, gosec.NewConfig()},
}
