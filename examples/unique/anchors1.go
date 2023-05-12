package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
)

var yamlCode = `
// START YAML OMIT
name: &name Carson
accounts:
  - id: 1
    name: *name
  - {id: 2, name: *name}
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
