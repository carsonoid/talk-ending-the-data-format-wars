package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
)

var tomlCode = `
// START TOML OMIT
"my string"
// END TOML OMIT
`

var hclCode = `
// START HCL OMIT
"my string"
// END HCL OMIT
`

func main() {
	run.Print(run.TOMLExample(tomlCode))
	run.Print(run.HCLExample(hclCode))
}

// START RUN OMIT
// Run it!
// END RUN OMIT
