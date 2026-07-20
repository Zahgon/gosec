package sarif

type Level string

const (
	None = Level("none")

	Note = Level("note")

	Warning = Level("warning")

	Error = Level("error")

	Version = "2.1.0"

	Schema = "https://raw.githubusercontent.com/oasis-tcs/sarif-spec/main/sarif-2.1/schema/sarif-schema-2.1.0.json"
)
