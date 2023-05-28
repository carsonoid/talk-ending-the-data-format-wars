package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
)

// //go:embed example.json
// var jsonBytes []byte

var jsonCode = `
// START JSON OMIT
{"date":"1979-05-27T07:32:00Z"}
// END JSON OMIT
`

var yamlCode = `
// START YAML OMIT
date1: 1979-05-27
date2: 1979-05-27T07:32:00Z
date3: 1979-05-27T00:32:00-07:00
// END YAML OMIT
`

var tomlCode = `
// START TOML OMIT
date1 = 1979-05-27
date2 = 1979-05-27T07:32:00Z
date3 = 1979-05-27T00:32:00-07:00
date4 = 1979-05-27 00:32:00-07:00
// END TOML OMIT
`

var hclCode = `
// START HCL OMIT
date = "1979-05-27T07:32:00Z"
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
