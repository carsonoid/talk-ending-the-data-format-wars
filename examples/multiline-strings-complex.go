package main

import (
	"bytes"

	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
)

const target = `
// START TARGET OMIT
line1
  line2
    line3
	< tab indented
// END TARGET OMIT
`

var jsonCode = `
// START JSON OMIT
{"key1":"line1\n  line2\n    line3\n\t< tab indented"}
// END JSON OMIT
`

var yamlCode = `
// START YAML OMIT
key1: "line1\n  line2\n    line3\n\t< tab indented"
key2: |-
  line1
    line2
      line3
  	< tab indented
// END YAML OMIT
`

var tomlCode = `
// START TOML OMIT
key1 = "line1\n  line2\n    line3\n\t< tab indented"
key2 = """
line1
  line2
    line3
	< tab indented"""
// END TOML OMIT
`

var hclCode = `
// START HCL OMIT
key1 = "line1\n  line2\n    line3\n\t< tab indented"
# again, heredocs only *almost* work so they are not shown here
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
