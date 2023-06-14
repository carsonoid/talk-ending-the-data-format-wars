package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
)

var yamlCode = `
// START YAML OMIT
# deduplicate env vars using anchors
x-base-env: &base-env
  FOO: BAR

services:
  first:
    image: my-image:latest
    environment: &first-env
      <<: *base-env
      FIZZ: BUZZ
  second:
    image: another-image:latest
    environment:
      <<: *first-env
      ZOT: QUIX
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
