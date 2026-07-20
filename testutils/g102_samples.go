package testutils

import "github.com/securego/gosec/v2"

var SampleCodeG102 = []CodeSample{

	{[]string{`
package main

import (
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"log"
	"net"
)

func main() {
		l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"log"
	"net"
)

func parseListenAddr(listenAddr string) (network string, addr string) {
	return "", ""
}

func main() {
	addr := ":2000"
	l, err := net.Listen(parseListenAddr(addr))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"log"
	"net"
)

const addr = ":2000"

func parseListenAddr(listenAddr string) (network string, addr string) {
	return "", ""
}

func main() {
	l, err := net.Listen(parseListenAddr(addr))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
}
`}, 1, gosec.NewConfig()},
	{[]string{`
package main

import (
	"log"
	"net"
)

const addr = "0.0.0.0:2000"

func main() {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
}
`}, 1, gosec.NewConfig()},
}
