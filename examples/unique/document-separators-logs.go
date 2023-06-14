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
{ name: Carson, id: 35 }
---
{ name: Tami, id: 14 }
...
---
{ name: Raul, id: 15 }
---
{ name: Ella, id: 23 }
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
