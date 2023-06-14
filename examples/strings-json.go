package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
)

var jsonCode = `
// START JSON OMIT
{
  "str": "my string",
  "bool": "true",
  "num": "1",

  "unicode": "\u03B1\u03B2\u03B3",

  "quoted": "'single' \"double\" quotes"
}
// END JSON OMIT
`

func main() {
	run.PrintAll(
		run.JSONExample(jsonCode),
	)
}

// START RUN OMIT
// Run it!
// END RUN OMIT
