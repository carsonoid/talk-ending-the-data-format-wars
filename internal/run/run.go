package run

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/hashicorp/hcl"
	"github.com/sanity-io/litter"
	"gopkg.in/yaml.v3"
)

func init() {
	timeType := reflect.TypeOf(time.Time{})
	litter.Config.DumpFunc = func(v reflect.Value, w io.Writer) bool {
		if v.Type() != timeType {
			return false
		}

		t := v.Interface().(time.Time)
		fmt.Fprintf(w, `{%s}`, t.Format(time.RFC3339))
		return true
	}
}

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

func dumpResult(example Example, obj any) {
	out := litter.Sdump(obj)
	out = strings.ReplaceAll(out, "interface {}", "any")
	fmt.Println(out)
	if example.Note != "" {
		fmt.Println("* ", example.Note)
	}
}

type Example struct {
	Type string
	Data string
	Note string
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
	fmt.Println("====== " + strings.ToUpper(example.Type) + " ======")
	switch example.Type {
	case "json":
		var jsonObj any
		err := json.Unmarshal(sanitize(example.Data), &jsonObj)
		if err != nil {
			fmt.Println("Error: " + err.Error())
		} else {
			dumpResult(example, jsonObj)
		}
	case "yaml":
		var yamlObj any
		err := yaml.Unmarshal(sanitize(example.Data), &yamlObj)
		if err != nil {
			fmt.Println("Error: " + err.Error())
		} else {
			dumpResult(example, yamlObj)
		}
	case "toml":
		var tomlObj any
		err := toml.Unmarshal(sanitize(example.Data), &tomlObj)
		if err != nil {
			fmt.Println("Error: " + err.Error())
		} else {
			dumpResult(example, tomlObj)
		}

	case "hcl":
		var hclObj any
		err := hcl.Unmarshal(sanitize(example.Data), &hclObj)
		if err != nil {
			fmt.Println("Error: " + err.Error())
		} else {
			dumpResult(example, hclObj)
		}
	default:
		panic("unknown type " + example.Type + " " + example.Data)
	}
}
