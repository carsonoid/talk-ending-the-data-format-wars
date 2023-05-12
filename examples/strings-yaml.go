package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
)

var yamlCode = `
// START YAML OMIT
str1: my string # no quotes required for many string values
str2: "my string"
str3: 'my string'
str4: !!str my string
str5: >-
  my string

bool1: 'true'
bool2: "true"
bool3: !!str true
bool4: >-
  true

unicode: "\u03B1\u03B2\u03B3"

complex1: '''single'' "double" quotes' # Note the '', not \'
complex2: "'single' \"double\" quotes" # supports \ sequences
complex3: >-
  'single' "double" quotes
// END YAML OMIT
`

func main() {
	run.PrintAll(
		run.YAMLExample(yamlCode),
	)
}

// START RUN OMIT
// Run it!
// END RUN OMIT
