package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
)

var jsonCode = `
// START JSON OMIT
{}
// END JSON OMIT
`

var yamlCode = `
// START YAML OMIT
# comment on a line
str1: "value" # comments can come after
str2: value # comments can come after
str3: "use quotes to keep # comments that come after"
str4: >-
  use a block scalar to keep # comments that come after

bool: true # comments can come after
num: 1 # comments can come after
// END YAML OMIT
`

var tomlCode = `
// START TOML OMIT
# key is the key
key1 = "value"
key2 = 'value' # comments can come after
key3 = true # comments can come after
// END TOML OMIT
`

var hclCode = `
// START HCL OMIT
# key is the key
key1 = "value"
key2 = "value" # comments can come after
// END HCL OMIT
`

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
