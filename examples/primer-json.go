package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
)

// //go:embed example.json
// var jsonBytes []byte

var jsonCode = `
// START JSON OMIT
{
    "str": "my string",
    "bool": "my bool",
    "map": {
        "true": true,
        "num": false
    },
	"array": [1,2,3]
}
// END JSON OMIT
`

var jsonCode2 = `
// START JSON2 OMIT
{"str":"my string","bool":"my bool","map":{"true":true,"num":false},"array":[1,2,3]}
// END JSON2 OMIT
`

func main() {
	run.PrintAll(
		run.JSONExample(jsonCode),
	)
}

// START RUN OMIT
// Run it!
// END RUN OMIT
