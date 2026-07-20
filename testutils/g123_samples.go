package testutils

import "github.com/securego/gosec/v2"

var SampleCodeG123 = []CodeSample{

	{[]string{`
package main

import (
	"crypto/tls"
	"crypto/x509"
)

func main() {
	_ = &tls.Config{
		VerifyPeerCertificate: func(_ [][]byte, _ [][]*x509.Certificate) error { return nil },
	}
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"crypto/tls"
	"crypto/x509"
)

func main() {
	_ = &tls.Config{
		GetConfigForClient: func(ch *tls.ClientHelloInfo) (*tls.Config, error) {
			_ = ch
			return &tls.Config{
				VerifyPeerCertificate: func(_ [][]byte, _ [][]*x509.Certificate) error { return nil },
			}, nil
		},
	}
}
`}, 2, gosec.NewConfig()},

	{[]string{`
package main

import (
	"crypto/tls"
	"crypto/x509"
)

func main() {
	_ = &tls.Config{
		VerifyPeerCertificate: func(_ [][]byte, _ [][]*x509.Certificate) error { return nil },
		VerifyConnection:      func(_ tls.ConnectionState) error { return nil },
	}
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import (
	"crypto/tls"
	"crypto/x509"
)

func main() {
	cfg := &tls.Config{}
	cfg.VerifyPeerCertificate = func(_ [][]byte, _ [][]*x509.Certificate) error { return nil }
	cfg.SessionTicketsDisabled = true
	_ = cfg
}
`}, 0, gosec.NewConfig()},
}
