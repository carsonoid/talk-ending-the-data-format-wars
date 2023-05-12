package main

import (
	// import embed for effect to enable go:embed
	"bytes"
	_ "embed"
	"encoding/json"

	"github.com/hashicorp/hcl"
	"github.com/hydronica/toml"
	"github.com/sanity-io/litter"
	"gopkg.in/yaml.v2"
)

// //go:embed example.json
// var jsonBytes []byte

var jsonBytes = sanitize([]byte(`
// START JSON OMIT
{"key1":"my string"}
// END JSON OMIT
`))

var yamlBytes = sanitize([]byte(`
// START YAML OMIT
key1: my string
key2: 'my string'
key3: "my string"
// END YAML OMIT
`))

var tomlBytes = sanitize([]byte(`
// START TOML OMIT
key1 = "my string"
key2 = 'my string'
// END TOML OMIT
`))

var hclBytes = sanitize([]byte(`
// START HCL OMIT
key1 = "my string"
// END HCL OMIT
`))

// //go:embed example.yaml
// var yamlBytes []byte

// //go:embed example.toml
// var tomlBytes []byte

// //go:embed example.hcl
// var hclBytes []byte

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
	jsonObj := make(map[string]string)
	err := json.Unmarshal(jsonBytes, &jsonObj)
	if err != nil {
		panic(err)
	}
	litter.Dump(jsonObj)

	yamlObj := make(map[string]string)
	err = yaml.Unmarshal(yamlBytes, &yamlObj)
	if err != nil {
		panic(err)
	}
	litter.Dump(yamlObj)

	tomlObj := make(map[string]string)
	err = toml.Unmarshal(tomlBytes, &tomlObj)
	if err != nil {
		panic(err)
	}
	litter.Dump(tomlObj)

	hclObj := make(map[string]string)
	err = hcl.Unmarshal(hclBytes, &hclObj)
	if err != nil {
		panic(err)
	}
	litter.Dump(hclObj)
}

// START RUN OMIT
// Run it!
// END RUN OMIT
