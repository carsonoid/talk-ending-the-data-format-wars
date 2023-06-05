package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
)

var tomlCode = `
// START TOML OMIT
[[people]]
name = "Alice"
age = 30
verified = true

[[people]]
name = "Bob"
age = 31
verified = false
// END TOML OMIT
`

var hclCode = `
// START HCL OMIT
people = [
	{ name = "Alice", age = 30, verified = true },
	{
		name = "Bob",
		age = 31,
		verified = false,
	},
]
// END HCL OMIT
`

func main() {
	run.PrintAll(
		run.TOMLExample(tomlCode),
		run.HCLExample(hclCode),
	)
}

// START RUN OMIT
// Run it!
// END RUN OMIT
