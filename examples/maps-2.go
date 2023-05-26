package main

import (
	"bytes"

	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
)

var jsonCode = `
// START JSON OMIT
{
  "myObject": {
	"key1": "value1",
	"key 2": "value2"
  }
}
// END JSON OMIT
`

var yamlCode = `
// START YAML OMIT
myObject:
  key1: value1
  "key 2": value2
// END YAML OMIT
`

var tomlCode = `
// START TOML OMIT
[myObject]
key1 = "value1"
"key 2" = "value2"
// END TOML OMIT
`

var hclCode = `
// START HCL OMIT
myObject {
  key1 = "value1"
  "key 2" = "value2\n"
}
// END HCL OMIT
`

// sanitize removes all lines that end in OMIT
func sanitize(b []byte) []byte {
	var out []byte
	for _, line := range bytes.Split(b, []byte("\n")) {
		if !bytes.HasSuffix(line, []byte("OMIT")) {
			out = append(out, line...)
			out = append(out, []byte("\n")...)
		}
	}
	return out
}
func main() {
	run.PrintAll(
		run.JSONExample(jsonCode),
		run.YAMLExample(yamlCode),
		run.TOMLExample(tomlCode),
		run.HCLExample(hclCode),
	)
}

// START RUN OMIT
// Run it!
// END RUN OMIT