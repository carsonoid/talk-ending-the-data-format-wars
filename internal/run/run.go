package run

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/hcl"
	"github.com/hydronica/toml"
	"github.com/sanity-io/litter"
	"gopkg.in/yaml.v2"
)

// sanitize removes all lines that end in OMIT
func sanitize(s string) []byte {
	b := []byte(s)
	var out []byte
	for _, line := range bytes.Split(b, []byte("\n")) {
		if !bytes.HasSuffix(line, []byte("OMIT")) {
			out = append(out, line...)
			out = append(out, []byte("\n")...)
		}
	}
	return out
}

func dumpResult(t string, obj any) {
	fmt.Println("====== " + strings.ToUpper(t) + " ======")
	litter.Dump(obj)
}

type Example struct {
	Type string
	Data string
}

func JSONExample(s string) Example {
	return Example{
		Type: "json",
		Data: s,
	}
}

func YAMLExample(s string) Example {
	return Example{
		Type: "yaml",
		Data: s,
	}
}

func TOMLExample(s string) Example {
	return Example{
		Type: "toml",
		Data: s,
	}
}

func HCLExample(s string) Example {
	return Example{
		Type: "hcl",
		Data: s,
	}
}

func PrintAll(examples ...Example) {
	for _, example := range examples {
		Print(example)
	}
}

func Print(example Example) {
	switch example.Type {
	case "json":
		jsonObj := make(map[string]string)
		err := json.Unmarshal(sanitize(example.Data), &jsonObj)
		if err != nil {
			panic(err)
		}
		dumpResult(example.Type, jsonObj)
	case "yaml":
		yamlObj := make(map[string]string)
		err := yaml.Unmarshal(sanitize(example.Data), &yamlObj)
		if err != nil {
			panic(err)
		}
		dumpResult(example.Type, yamlObj)
	case "toml":
		tomlObj := make(map[string]string)
		err := toml.Unmarshal(sanitize(example.Data), &tomlObj)
		if err != nil {
			panic(err)
		}
		dumpResult(example.Type, tomlObj)

	case "hcl":
		hclObj := make(map[string]string)
		err := hcl.Unmarshal(sanitize(example.Data), &hclObj)
		if err != nil {
			panic(err)
		}
		dumpResult(example.Type, hclObj)
	default:
		panic("unknown type " + example.Type + " " + example.Data)
	}
}
