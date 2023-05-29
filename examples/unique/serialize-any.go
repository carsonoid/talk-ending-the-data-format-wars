package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
)

var jsonCode = `
// START JSON OMIT
"my string"
// END JSON OMIT
`

var jsonCode2 = `
// START JSON2 OMIT
1
// END JSON2 OMIT
`

var jsonCode3 = `
// START JSON3 OMIT
true
// END JSON3 OMIT
`

var yamlCode = `
// START YAML OMIT
my string
// END YAML OMIT
`

var yamlCode2 = `
// START YAML2 OMIT
1
// END YAML2 OMIT
`

var yamlCode3 = `
// START YAML3 OMIT
true
// END YAML3 OMIT
`

var yamlCode4 = `
// START YAML4 OMIT
# empty is valid in YAML but not JSON
// END YAML4 OMIT
`

func main() {
	run.Print(run.JSONExample(jsonCode))
	run.Print(run.JSONExample(jsonCode2))
	run.Print(run.JSONExample(jsonCode3))
	run.Print(run.YAMLExample(yamlCode))
	run.Print(run.YAMLExample(yamlCode2))
	run.Print(run.YAMLExample(yamlCode3))
	run.Print(run.YAMLExample(yamlCode4))
}

// START RUN OMIT
// Run it!
// END RUN OMIT
