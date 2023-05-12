package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
)

var yamlCode = `
// START YAML OMIT
str: my string
map:
  "true": true
  !!str false: false
array:
  - 1
  - 2
  - 3
jsonExample: {"str":"my string","map":{"true":true,"num":false},"array":[1,2,3]}
// END YAML OMIT
`

func main() {
	run.PrintAll(
		run.YAMLExample(yamlCode),
	)
}

// START RUN OMIT
// Run it!
// END RUN OMIT
