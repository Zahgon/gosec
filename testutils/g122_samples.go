package testutils

import "github.com/securego/gosec/v2"

var SampleCodeG122 = []CodeSample{

	{[]string{`
package main

import (
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	_ = filepath.WalkDir("/tmp", func(path string, d fs.DirEntry, err error) error {
		_ = d
		_ = err
		return os.Remove(path)
	})
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	_ = filepath.WalkDir("/var/data", func(path string, d fs.DirEntry, err error) error {
		_ = d
		_ = err
		target := path + ".bak"
		_, openErr := os.OpenFile(target, os.O_RDWR|os.O_CREATE, 0o600)
		return openErr
	})
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"io/fs"
	"path/filepath"
)

func main() {
	_ = filepath.WalkDir("/tmp", func(path string, d fs.DirEntry, err error) error {
		_ = path
		_ = d
		_ = err
		return nil
	})
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import (
	"os"
	"path/filepath"
)

func main() {
	_ = filepath.Walk("/tmp", func(path string, info os.FileInfo, err error) error {
		_ = path
		_ = info
		_ = err
		return os.Remove("/tmp/fixed-file")
	})
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import (
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	root, err := os.OpenRoot("/tmp")
	if err != nil {
		return
	}
	defer root.Close()

	_ = filepath.WalkDir("/tmp", func(path string, d fs.DirEntry, err error) error {
		_ = d
		_ = err
		_, openErr := root.Open(path)
		return openErr
	})
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import (
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	root, err := os.OpenRoot("/tmp")
	if err != nil {
		return
	}
	defer root.Close()

	_ = filepath.WalkDir("/tmp", func(path string, d fs.DirEntry, err error) error {
		_ = d
		_ = err
		return root.Remove(path)
	})
}
`}, 0, gosec.NewConfig()},
}
