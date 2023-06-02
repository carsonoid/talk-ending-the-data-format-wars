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
json = <<EOJSON
{ type = "json", data = { str = "\"Hello World\"", bool = true } }
EOJSON

yaml = <<EOYAML
type: yaml
data:
  str: '"Hello World"'
  bool: true
EOYAML

toml = <<EOTOML
type = "toml"
[[data]]
str = '"Hello World"'
bool = true
EOTOML

hcl = <<EOHCL
type = "hcl"
data = {
  str = "\"Hello World\"",
  bool = true,
}
EOHCL
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
