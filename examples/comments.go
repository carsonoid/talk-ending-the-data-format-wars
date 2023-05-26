package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
)

var jsonCode = `
// START JSON OMIT
{}
// END JSON OMIT
`

var yamlCode = `
// START YAML OMIT
# key is the key
key: value
// END YAML OMIT
`

var tomlCode = `
// START TOML OMIT
# key is the key
key = "value"
// END TOML OMIT
`

var hclCode = `
// START HCL OMIT
# key is the key
key = "value"
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
