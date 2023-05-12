package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
)

// //go:embed example.json
// var jsonBytes []byte

var tomlCode = `
// START TOML OMIT
bools = [ true, false, true ]
numbers = [ 1, 2, 3, 4, 5 ]
floats = [ 1.1, 2.2, 3.3   ] # whitespace doesn't matter
names = [ # can span multiple lines
  "Alice",
  "Bob",
  "Carol", # trailing comma is ok
]
// END TOML OMIT
`

var hclCode = `
// START HCL OMIT
bools = [ true, false, true ]
numbers = [ 1, 2, 3, 4, 5 ]
floats = [ 1.1, 2.2, 3.3   ] # whitespace doesn't matter
names = [ # can span multiple lines
  "Alice",
  "Bob",
  "Carol", # trailing comma is ok
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
