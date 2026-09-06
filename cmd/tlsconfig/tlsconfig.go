package main

import (
	"bytes"
	"flag"
	"go/format"
	"log"
	"os"
	"path/filepath"
)

var (
	pkg        = flag.String("pkg", "rules", "package name to be added to the output file")
	outputFile = flag.String("outputFile", "tls_config.go", "name of the output file")
)

const TLSConfURL = "https://statics.tls.security.mozilla.org/server-side-tls-conf.json"

type ServerSideTLSJson struct {
	Configurations map[string]Configuration `json:"configurations"`
	Version        float64                  `json:"version"`
}

type Configuration struct {
	OpenSSLCiphersuites   []string `json:"openssl_ciphersuites"`
	OpenSSLCiphers        []string `json:"openssl_ciphers"`
	TLSVersions           []string `json:"tls_versions"`
	TLSCurves             []string `json:"tls_curves"`
	CertificateTypes      []string `json:"certificate_types"`
	CertificateCurves     []string `json:"certificate_curves"`
	CertificateSignatures []string `json:"certificate_signatures"`
	RsaKeySize            float64  `json:"rsa_key_size"`
	DHParamSize           float64  `json:"dh_param_size"`
	ECDHParamSize         float64  `json:"ecdh_param_size"`
	HstsMinAge            float64  `json:"hsts_min_age"`
	OldestClients         []string `json:"oldest_clients"`
	OCSPStaple            bool     `json:"ocsp_staple"`
	ServerPreferredOrder  bool     `json:"server_preferred_order"`
	MaxCertLifespan       float64  `json:"maximum_certificate_lifespan"`
}

type goCipherConfiguration struct {
	Name       string
	Ciphers    []string
	MinVersion string
	MaxVersion string
}

type goTLSConfiguration struct {
	cipherConfigs []goCipherConfiguration
}

func getTLSConfFromURL(url string) (*ServerSideTLSJson, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getGoCipherConfig(name string, sstls ServerSideTLSJson) (goCipherConfiguration, error) {
	_ = "STUB: not implemented"
	return *new(goCipherConfiguration), nil
}

func getGoTLSConf() (goTLSConfiguration, error) {
	_ = "STUB: not implemented"
	return *new(goTLSConfiguration), nil
}

func getCurrentDir() (string, error) { _ = "STUB: not implemented"; return "", nil }

func main() {
	dir, err := getCurrentDir()
	if err != nil {
		log.Fatalln(err)
	}
	tlsConfig, err := getGoTLSConf()
	if err != nil {
		log.Fatalln(err)
	}

	var buf bytes.Buffer
	err = generatedHeaderTmpl.Execute(&buf, *pkg)
	if err != nil {
		log.Fatalf("Failed to generate the header: %v", err)
	}
	for _, cipherConfig := range tlsConfig.cipherConfigs {
		err := generatedRuleTmpl.Execute(&buf, cipherConfig)
		if err != nil {
			log.Fatalf("Failed to generated the cipher config: %v", err)
		}
	}

	src, err := format.Source(buf.Bytes())
	if err != nil {
		log.Printf("warnings: Failed to format the code: %v", err)
		src = buf.Bytes()
	}

	outputPath := filepath.Join(dir, *outputFile)
	if err := os.WriteFile(outputPath, src, 0o644); err != nil {
		log.Fatalf("Writing output: %s", err)
	}
}
