package main

import (
	"github.com/carsonoid/talk-ending-the-data-format-wars/internal/run"
)

var tomlCode = `
// START TOML OMIT
str1 = "my string"
str2 = 'my string'
str3 = '''my string'''
str4 = """my string"""

bool1 = "true"
bool2 = 'true'
bool3 = '''true'''
bool4 = """true"""

unicode = "\u03B1\u03B2\u03B3"

complex1 = "'single' \"double\" quotes"
# NO escaping in single quotes in TOML
# complex2 = '''single'' "double" quotes'
complex3 = """'single' "double" quotes"""
complex4 = ''''single' "double" quotes'''
// END TOML OMIT
`

func main() {
	run.PrintAll(
		run.TOMLExample(tomlCode),
	)
}

// START RUN OMIT
// Run it!
// END RUN OMIT
