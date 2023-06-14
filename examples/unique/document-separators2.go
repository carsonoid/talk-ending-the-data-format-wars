package main

import (
	"bytes"
	"fmt"

	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
	"gopkg.in/yaml.v3"
)

var yamlCode = `
// START YAML OMIT
---
name: Carson
id: 35
---
name: Tami
id: 14

# Note the ... to say
# "end of stream" rather than "start of document"
#
# Without this, the stream wouldn't be able to tell
# the difference between missing data and an and of data
...
// END YAML OMIT
`

func main() {
	ex := run.YAMLExample(string(run.Sanitize(yamlCode)))
	dec := yaml.NewDecoder(bytes.NewBufferString(ex.Data))
	for {
		var m map[string]any
		err := dec.Decode(&m)
		if err != nil {
			fmt.Println(err)
			break
		}
		run.DumpResult(ex, m)
	}
}

// START RUN OMIT
// Run it!
// END RUN OMIT
