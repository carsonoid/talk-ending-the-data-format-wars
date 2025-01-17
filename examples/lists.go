package main

import (
	"bytes"

	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
)

var jsonCode = `
// START JSON OMIT
{ "myList": [ "value", false ] }
// END JSON OMIT
`

var yamlCode = `
// START YAML OMIT
myList1: [ "value", false ]
myList2: [
	value, false,
]
myList3:
  - value
  - !!bool false # type tag not required but valid
// END YAML OMIT
`

var tomlCode = `
// START TOML OMIT
myList1 = ["value",false]
myList2 = [
  "value",
  false,
]
// END TOML OMIT
`

var hclCode = `
// START HCL OMIT
myList1 = [ "value", false ]
myList2 = [
  "value1",
  false,
]
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
