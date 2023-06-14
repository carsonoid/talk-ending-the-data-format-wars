package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
)

// {
//   "type": "json",
//   "data": { "str": "\"Hello World\"", "bool": true }
// }

var tomlCode = `
// START TOML OMIT
json = '{"type":"json","data":{"str":"\"Hello World\"","bool":true}}'

yaml = '''
type: yaml
data:
  str: '"Hello World"'
  bool: true'''

toml = '''
type = "toml"
[[data]]
str = '"Hello World"'
bool = true'''

hcl = '''
type = "hcl"
data = {
  str = "Hello World",
  bool = true,
}'''
// END TOML OMIT
`

func main() {
	run.PrintNested(
		run.TOMLExample(tomlCode),
	)
}

// START RUN OMIT
// Run it!
// END RUN OMIT
