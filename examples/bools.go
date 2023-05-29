package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
)

// //go:embed example.json
// var jsonBytes []byte

var jsonCode = `
// START JSON OMIT
{
  "true": true,
  "false": false
}
// END JSON OMIT
`

var yamlCode = `
// START YAML OMIT
"true": true
true2: yes
true3: Y
!!str false: false
false2: no
false3: N
// END YAML OMIT
`

var tomlCode = `
// START TOML OMIT
true = true
false = false
// END TOML OMIT
`

var hclCode = `
// START HCL OMIT
"true" = true
"false" = false
// END HCL OMIT
`

func main() {
	run.PrintAs(map[string]bool{}, run.JSONExample(jsonCode))
	run.PrintAs(map[string]bool{}, run.YAMLExample(yamlCode))
	run.PrintAs(map[string]bool{}, run.TOMLExample(tomlCode))
	run.PrintAs(map[string]bool{}, run.HCLExample(hclCode))
}

// START RUN OMIT
// Run it!
// END RUN OMIT
