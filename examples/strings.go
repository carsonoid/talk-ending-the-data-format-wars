package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
)

// //go:embed example.json
// var jsonBytes []byte

var jsonCode = `
// START JSON OMIT
{
  "key1":"my string",
  "key2":"\"quoted\" string",
  "key3": "true"
}
// END JSON OMIT
`

var yamlCode = `
// START YAML OMIT
key1: my string
key2: '"quoted" string'
key3: !!str true
// END YAML OMIT
`

var tomlCode = `
// START TOML OMIT
key1 = "my string"
key2 = '"quoted" string'
key3 = "true"
// END TOML OMIT
`

var hclCode = `
// START HCL OMIT
key1 = "my string"
key2 = "\"quoted\" string"
key3 = "true"
// END HCL OMIT
`

func main() {
	run.PrintAll(
		run.JSONExample(jsonCode),
		run.YAMLExample(yamlCode),
		run.TOMLExample(tomlCode),
		run.HCLExample(hclCode),
	)
}

// START RUN OMIT
// Run it!
// END RUN OMIT
