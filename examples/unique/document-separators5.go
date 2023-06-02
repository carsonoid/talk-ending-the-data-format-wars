package main

import (
	"bytes"
	"fmt"

	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
	"gopkg.in/yaml.v3"
)

const user1 = `
// START USER1 OMIT
{ name: Carson, id: 35 }
// END USER1 OMIT
`

const user2 = `
// START USER2 OMIT
{ name: Tami, id: 14 }
// END USER2 OMIT
`

const user3 = `
// START USER3 OMIT
name: Raul
id: 15
// END USER3 OMIT
`

const bash = `
// START BASH OMIT
echo -e "\n---\n" > sep.yaml
cat user1.yaml sep.yaml user2.yaml sep.yaml user3.yaml | kubectl apply -f -
// END BASH OMIT
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
