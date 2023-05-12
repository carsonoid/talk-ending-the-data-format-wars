package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
)

var hclCode = `
// START HCL OMIT
str = "my string"
map = {
	"true" = true
	"false" = false
}
array = [1, 2, 3]
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
