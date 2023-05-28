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
	"github.com/mitchellh/go-wordwrap"
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

func PrintAllAs[T any](dest T, examples ...Example) {
	for _, example := range examples {
		PrintAs(dest, example)
	}
}

func Print(example Example) {
	PrintAs(map[string]any{}, example)
}

func PrintErr(err error) {
	errStr := err.Error()
	errStr = wordwrap.WrapString(errStr, 60)
	fmt.Println("Error: " + errStr)
}

func PrintAs[T any](dest T, example Example) {
	fmt.Println("====== " + strings.ToUpper(example.Type) + " ======")
	switch example.Type {
	case "json":
		err := json.Unmarshal(sanitize(example.Data), &dest)
		if err != nil {
			PrintErr(err)
		} else {
			dumpResult(example, dest)
		}
	case "yaml":
		err := yaml.Unmarshal(sanitize(example.Data), &dest)
		if err != nil {
			PrintErr(err)
		} else {
			dumpResult(example, dest)
		}
	case "toml":
		err := toml.Unmarshal(sanitize(example.Data), &dest)
		if err != nil {
			PrintErr(err)
		} else {
			dumpResult(example, dest)
		}

	case "hcl":
		err := hcl.Unmarshal(sanitize(example.Data), &dest)
		if err != nil {
			PrintErr(err)
		} else {
			dumpResult(example, dest)
		}
	default:
		panic("unknown type " + example.Type + " " + example.Data)
	}
}
