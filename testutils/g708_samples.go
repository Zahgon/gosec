package testutils

import "github.com/securego/gosec/v2"

var SampleCodeG708 = []CodeSample{

	{[]string{`
package main

import (
	"net/http"
	"text/template"
)

func handler(w http.ResponseWriter, r *http.Request) {
	userTmpl := r.URL.Query().Get("tmpl")
	t, _ := template.New("page").Parse(userTmpl)
	t.Execute(w, nil)
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.New("page").Parse(` + "`<h1>Hello {{.}}</h1>`" + `))

func handler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	tmpl.Execute(w, name)
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.New("").Parse(` + "`{{define \"greeting\"}}Hello {{.}}{{end}}`" + `))

func handler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	tmpl.ExecuteTemplate(w, "greeting", name)
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.New("page").Parse(` + "`<h1>Hello {{.}}</h1>`" + `))

func handler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	tmpl.Execute(w, name)
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import (
	"os"
	"text/template"
)

func main() {
	tmpl := template.Must(template.New("page").Parse(` + "`Hello {{.}}`" + `))
	tmpl.Execute(os.Stdout, "World")
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import (
	"html"
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.New("page").Parse(` + "`<h1>Hello {{.}}</h1>`" + `))

func handler(w http.ResponseWriter, r *http.Request) {
	safe := html.EscapeString(r.FormValue("name"))
	tmpl.Execute(w, safe)
}
`}, 0, gosec.NewConfig()},
}
