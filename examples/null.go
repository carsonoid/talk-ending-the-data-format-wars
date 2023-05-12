package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
)

// //go:embed example.json
// var jsonBytes []byte

var jsonCode = `
// START JSON OMIT
{"key":null}
// END JSON OMIT
`

var yamlCode = `
// START YAML OMIT
key1: null
key2: Null
key3: NULL
key4: ~
key5:
// END YAML OMIT
`

var tomlCode = `
// START TOML OMIT
key = null
// END TOML OMIT
`

var hclCode = `
// START HCL OMIT
key = null
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
