package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
)

var hclCode = `
// START HCL OMIT
str = "my string"
bool = "true"
num = "1"

unicode = "\u03B1\u03B2\u03B3"

quoted = "'single' \"double\" quotes"
// END HCL OMIT
`

func main() {
	run.PrintAll(
		run.HCLExample(hclCode),
	)
}

// START RUN OMIT
// Run it!
// END RUN OMIT
