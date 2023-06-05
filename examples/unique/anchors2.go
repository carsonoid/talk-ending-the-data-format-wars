package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
)

var yamlCode = `
// START YAML OMIT
name: &name Carson
users:
  - &carson
    id: 35
    name: *name
accounts:
  - &account1
    user: *carson
    balance: 100
  - user: *carson
    balance: 200
defaultAccount: *account1
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
