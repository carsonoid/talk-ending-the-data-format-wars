package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
)

// //go:embed example.json
// var jsonBytes []byte

var jsonCode = `
// START JSON OMIT
{
  "bools": [ true, false, true ],
  "numbers": [ 1, 2, 3, 4, 5 ],
  "floats": [ 1.1, 2.2, 3.3 ],
  "names": [ "Alice", "Bob", "Carol" ]
}
// END JSON OMIT
`

var yamlCode = `
// START YAML OMIT
bools: [ true, false, true ]
numbers:
- 1
- 2
- 3
floats:
  - 1.1
  - 2.2
  - 3.3
names:
    - Alice
    - Bob
    - Carol
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
