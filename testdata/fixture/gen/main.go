// Command gen is the fixture's generator, invoked via the //go:generate
// directive in message.go. It renders an embedded template into a
// deterministic generated file so the action's self-test can assert both
// the pass and fail paths.
package main

import (
	_ "embed"
	"os"
	"text/template"
)

//go:embed message_generated.go.tmpl
var tmpl string

type data struct {
	Value string
}

func main() {
	t := template.Must(template.New("generated").Parse(tmpl))

	f, err := os.Create("message_generated.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := t.Execute(f, data{Value: "generated-value"}); err != nil {
		panic(err)
	}
}
