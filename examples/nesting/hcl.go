package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
)

// {
//   "type": "json",
//   "data": { "str": "\"Hello World\"", "bool": true }
// }

var hclCode = `
// START HCL OMIT
json = "{\"type\":\"json\",\"data\":{\"str\":\"\\\"Hello World\\\"\",\"bool\": true}}"
yaml = "type: yaml\ndata:\n  str: '\"Hello World\"'\n  bool: true"
toml = "type = \"toml\"\n[[data]]\nstr = '\"Hello World\"'\nbool = true"
hcl = "type = \"hcl\"\ndata = {\n  str = \"\\\"Hello World\\\"\",\n  bool = true}"
// END HCL OMIT
`

func main() {
	run.PrintNested(
		run.HCLExample(hclCode),
	)
}

// START RUN OMIT
// Run it!
// END RUN OMIT
