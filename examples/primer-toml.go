package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
)

var tomlCode = `
// START TOML OMIT
str = "my string"
bool = "my bool"
map.true = true
map.num = false
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
