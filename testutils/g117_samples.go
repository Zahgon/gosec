package testutils

import "github.com/securego/gosec/v2"

var SampleCodeG117 = []CodeSample{

	{[]string{`
package main

import "encoding/json"

type Config struct {
	Password string
}

func main() {
	_, _ = json.Marshal(Config{})
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

type Config struct {
	APIKey *string ` + "`json:\"api_key\"`" + `
}

func main() {
	_, _ = json.MarshalIndent(Config{}, "", "  ")
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	PrivateKey []byte ` + "`json:\"private_key\"`" + `
}

func main() {
	_ = json.NewEncoder(os.Stdout).Encode(Config{})
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

type Config struct {
	Password string ` + "`json:\"text_field\"`" + `
}

func main() {
	_, _ = json.Marshal(Config{})
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

type Config struct {
	SafeField string ` + "`json:\"api_key\"`" + `
}

func main() {
	_, _ = json.Marshal(Config{})
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

type Config struct {
	Token string ` + "`json:\"auth_token\"`" + `
}

func main() {
	_, _ = json.Marshal(Config{})
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

type Config struct {
	Key string ` + "`json:\"access-key\"`" + `
}

func main() {
	_, _ = json.Marshal(Config{})
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

type Config struct {
	Secret string ` + "`json:\",omitempty\"`" + `
}

func main() {
	_, _ = json.Marshal(Config{})
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

type Config struct {
	ApiTokens []string
}

func main() {
	_, _ = json.Marshal(Config{})
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

type Config struct {
	RefreshTokens []string ` + "`json:\"refresh_tokens\"`" + `
}

func main() {
	_, _ = json.Marshal(Config{})
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

type Config struct {
	AccessTokens []*string
}

func main() {
	_, _ = json.Marshal(Config{})
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

type Config struct {
	CustomSecret string ` + "`json:\"my_custom_secret\"`" + `
}

func main() {
	_, _ = json.Marshal(Config{})
}
`}, 1, func() gosec.Config {
		cfg := gosec.NewConfig()
		cfg.Set("G117", map[string]interface{}{
			"pattern": "(?i)custom[_-]?secret",
		})
		return cfg
	}()},

	{[]string{`
package main

import "encoding/json"

type Config struct {
	Password string
}

func main() {
	_, _ = json.Marshal(&Config{})
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

type Config struct {
	Password string
}

func main() {
	_, _ = json.Marshal([]Config{{}})
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

type Config struct {
	Password string
}

func main() {
	_, _ = json.Marshal(map[string]Config{"x": {}})
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "go.yaml.in/yaml/v3"

type Config struct {
	Password string ` + "`yaml:\"password\"`" + `
}

func main() {
	_, _ = yaml.Marshal(Config{})
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/xml"

type Config struct {
	SafeField string ` + "`xml:\"api_key\"`" + `
}

func main() {
	_, _ = xml.Marshal(Config{})
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "github.com/BurntSushi/toml"
import "os"

type Config struct {
	Password string ` + "`toml:\"password\"`" + `
}

func main() {
	_ = toml.NewEncoder(os.Stdout).Encode(Config{})
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

type Config struct {
	Password string
}

func main() {}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import (
	"bytes"
	"text/template"
)

func main() {
	t := template.Must(template.New("x").Parse("{{.Username}}"))
	var tpl bytes.Buffer
	_ = t.Execute(&tpl, struct {
		Username string
		Password string
	}{})
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

type AppConfig struct {
	ApiSecret string ` + "`env:\"API_SECRET\"`" + `
}

func main() {}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

type Config struct {
	Password string ` + "`json:\"-\"`" + `
}

func main() {
	_, _ = json.Marshal(Config{})
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "go.yaml.in/yaml/v3"

type Config struct {
	Password string ` + "`yaml:\"-\"`" + `
}

func main() {
	_, _ = yaml.Marshal(Config{})
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/xml"

type Config struct {
	Password string ` + "`xml:\"-\"`" + `
}

func main() {
	_, _ = xml.Marshal(Config{})
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "github.com/BurntSushi/toml"
import "os"

type Config struct {
	Password string ` + "`toml:\"-\"`" + `
}

func main() {
	_ = toml.NewEncoder(os.Stdout).Encode(Config{})
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

type Config struct {
	UserID string ` + "`json:\"user_id\"`" + `
}

func main() {
	_, _ = json.Marshal(Config{})
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

func main() {
	_, _ = json.Marshal("api_key")
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

type Config struct {
	password string
}

func main() {
	_, _ = json.Marshal(Config{})
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

type Config struct {
	password string ` + "`json:\"password\"`" + `
}

func main() {
	_, _ = json.Marshal(Config{})
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

type Config struct {
	SafeField string ` + "`json:\"-,\"`" + `
}

func main() {
	_, _ = json.Marshal(Config{})
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

type Config struct {
	MaxTokens int
}

func main() {
	_, _ = json.Marshal(Config{})
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

type Config struct {
	RedactionTokens []string ` + "`json:\"redactionTokens,omitempty\"`" + `
}

func main() {
	_, _ = json.Marshal(Config{})
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

type Config struct {
	Safe, Password string
}

func main() {
	_, _ = json.Marshal(Config{})
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

type Config struct {
	Password string
}

func main() {
	_, _ = json.Marshal(Config{}) // #nosec G117
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

type Config struct {
	Password string
}

func main() {
	// #nosec G117 -- false positive
	_, _ = json.Marshal(Config{})
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

type Config struct {
	APIKey string ` + "`json:\"api_key\"`" + `
}

func main() {
	_, _ = json.Marshal(Config{}) // #nosec G117 -- public key
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

type Config struct {
	Password string
}

func main() {
	_, _ = json.MarshalIndent(Config{}, "", "  ") // #nosec G117
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Password string
}

func main() {
	_ = json.NewEncoder(os.Stdout).Encode(Config{}) // #nosec G117
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "go.yaml.in/yaml/v3"

type Config struct {
	Password string
}

func main() {
	_, _ = yaml.Marshal(Config{}) // #nosec G117
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/xml"

type Config struct {
	Password string
}

func main() {
	_, _ = xml.Marshal(Config{}) // #nosec G117
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "github.com/BurntSushi/toml"
import "os"

type Config struct {
	Password string
}

func main() {
	_ = toml.NewEncoder(os.Stdout).Encode(Config{}) // #nosec G117
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

type Credentials struct {
	Username string
	Password string ` + "`json:\"-\"`" + `
}

func (c Credentials) MarshalJSON() ([]byte, error) {
	type Aux struct {
		Username string
		Password string
	}
	return json.Marshal(Aux{
		Username: c.Username,
		Password: mask(c.Password),
	})
}

func mask(input string) string {
	return "****"
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

type Secret struct {
	Token string
}

func (s Secret) MarshalYAML() (interface{}, error) {
	type safe struct {
		Token string
	}
	b, err := json.Marshal(safe{Token: redact(s.Token)})
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

func redact(s string) string { return "***" }
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

type Credentials struct {
	Username string
	Password string
}

func (c Credentials) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

type Credentials struct {
	Username string
	Password string
}

func (c Credentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct{ Username string }{Username: c.Username})
}

func main() {
	_, _ = json.Marshal(Credentials{})
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

type Credentials struct {
	Username string
	Password string
}

func (c *Credentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct{ Username string }{Username: c.Username})
}

func main() {
	_, _ = json.Marshal(&Credentials{})
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

type Credentials struct {
	Username string
	Password string
}

func (c Credentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct{ Username string }{Username: c.Username})
}

func main() {
	_, _ = json.Marshal([]Credentials{{}})
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

type LogEntry struct {
	User     string
	Password string
}

func mask(s string) string { return "****" }

func main() {
	_, _ = json.Marshal(LogEntry{
		User:     "admin",
		Password: mask("secret123"),
	})
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

type LogEntry struct {
	User     string
	Password string
}

func mask(s string) string { return "****" }

func main() {
	_, _ = json.Marshal(&LogEntry{
		User:     "admin",
		Password: mask("secret123"),
	})
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

type LogEntry struct {
	User     string
	Password string
}

func main() {
	pw := "secret123"
	_, _ = json.Marshal(LogEntry{
		User:     "admin",
		Password: pw,
	})
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

import "encoding/json"

type Credentials struct {
	Username string
	Password string
}

type LogEntry struct {
	User     string
	Password string
}

func logCreds(c Credentials) {
	_, _ = json.Marshal(LogEntry{
		User:     c.Username,
		Password: c.Password,
	})
}
`}, 1, gosec.NewConfig()},

	{[]string{`
package main

type Config struct {
	Password string
}

func Marshal(any) {}

func main() {
	Marshal(Config{})
}
`}, 0, gosec.NewConfig()},

	{[]string{`
package main

type Encoder struct{}

func (Encoder) Encode(any) error { return nil }

type Config struct {
	Password string
}

func main() {
	_ = Encoder{}.Encode(Config{})
}
`}, 0, gosec.NewConfig()},
}
