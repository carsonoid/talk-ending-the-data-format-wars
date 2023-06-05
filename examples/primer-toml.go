package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
)

var tomlCode = `
// START TOML OMIT
str = "my string"
map.true = true
map.false = false
array = [1, 2, 3]
// END TOML OMIT
`

func main() {
	run.PrintAll(
		run.TOMLExample(tomlCode),
	)
}

// START RUN OMIT
// Run it!
// END RUN OMIT
