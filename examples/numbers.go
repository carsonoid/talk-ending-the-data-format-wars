package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
)

// //go:embed example.json
// var jsonBytes []byte

var jsonCode = `
// START JSON OMIT
{
  "number1":1,
  "number2":2.0,
  "number3":3.0e1,
  "number4":1000000
}
// END JSON OMIT
`

var yamlCode = `
// START YAML OMIT
number1: 1
number2: 2.0
number3: 3.0e1
number4: 1_000_000
// END YAML OMIT
`

var tomlCode = `
// START TOML OMIT
number1 = 1
number2 = 2.0
number3 = 3.0e1
number4 = 1_000_000
// END TOML OMIT
`

var hclCode = `
// START HCL OMIT
number1 = 1
number2 = 2.0
number3 = 3.0e1
number4 = 1000000
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
