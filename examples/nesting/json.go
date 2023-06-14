package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
)

var jsonCode = `
// START JSON OMIT
{
  "json":
    "{\"type\":\"json\",\"data\":{\"str\":\"\\\"Hello World\\\"\",\"bool\": true}}",
  "yaml":
    "type: yaml\ndata:\n  str: '\"Hello World\"'\n  bool: true",
  "toml":
    "type = \"toml\"\n[[data]]\nstr = '\"Hello World\"'\nbool = true",
  "hcl":
    "type = \"hcl\"\ndata = {\n  str = \"\\\"Hello World\\\"\",\n  bool = true}"
}
// END JSON OMIT
`

func main() {
	run.PrintNested(
		run.JSONExample(jsonCode),
	)
}

// START RUN OMIT
// Run it!
// END RUN OMIT
