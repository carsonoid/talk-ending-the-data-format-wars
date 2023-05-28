package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
)

// //go:embed example.json
// var jsonBytes []byte

var jsonCode = `
// START JSON OMIT
{
	"people": [
		{ "name": "Alice", "age": 30, "verified": true },
		{ "name": "Bob", "age": 31, "verified": false }
	]
}
// END JSON OMIT
`

var yamlCode = `
// START YAML OMIT
people:
  - name: Alice
    age: 30
    verified: true
  - name: Bob
    age: 31
    verified: false
// END YAML OMIT
`

func main() {
	run.PrintAll(
		run.JSONExample(jsonCode),
		run.YAMLExample(yamlCode),
	)
}

// START RUN OMIT
// Run it!
// END RUN OMIT
