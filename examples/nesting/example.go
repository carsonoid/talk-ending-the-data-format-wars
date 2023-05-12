package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
)

var jsonCode = `
// START JSON OMIT
{
  "type": "json",
  "data": { "str": "\"Hello World\"", "bool": true }
}
// END JSON OMIT
`

var yamlCode = `
// START YAML OMIT
type: yaml
data:
  str: '"Hello World"'
  bool: true
// END YAML OMIT
`

var tomlCode = `
// START TOML OMIT
type = "toml"
[[data]]
str = '"Hello World"'
bool = true
// END TOML OMIT
`

var hclCode = `
// START HCL OMIT
type = "hcl"
data = {
  str = "\"Hello World\"",
  bool = true
}
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
