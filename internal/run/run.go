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

// Sanitize removes all lines that end in OMIT
func Sanitize(s string) []byte {
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

func DumpResult(example Example, obj any) {
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

func PrintNested(example Example) {
	Print(example)

	var dest map[string]string
	err := Unmarshal(example, &dest)
	if err != nil {
		PrintErr(err)
		return
	}

	fmt.Println("\n\n------ " + "Parsed Nested Values" + "------")
	for _, k := range []string{"json", "yaml", "toml", "hcl"} {
		v, ok := dest[k]
		if !ok {
			continue
		}

		var example Example
		switch k {
		case "json":
			example = JSONExample(v)
		case "yaml":
			example = YAMLExample(v)
		case "toml":
			example = TOMLExample(v)
		case "hcl":
			example = HCLExample(v)
		default:
			fmt.Println("unknown type")
			continue
		}
		Print(example)
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
	var dest any
	PrintAs(dest, example)
}

func PrintErr(err error) {
	errStr := err.Error()
	errStr = wordwrap.WrapString(errStr, 70)
	fmt.Println("Error: " + errStr)
}

func PrintAs[T any](dest T, example Example) {
	fmt.Println("====== " + strings.ToUpper(example.Type) + " ======")
	err := Unmarshal(example, &dest)
	if err != nil {
		PrintErr(err)
		return
	}
	DumpResult(example, dest)
}

func Unmarshal(example Example, dest any) error {
	switch example.Type {
	case "json":
		err := json.Unmarshal(Sanitize(example.Data), dest)
		if err != nil {
			return err
		}
	case "yaml":
		err := yaml.Unmarshal(Sanitize(example.Data), dest)
		if err != nil {
			return err
		}
	case "toml":
		err := toml.Unmarshal(Sanitize(example.Data), dest)
		if err != nil {
			return err
		}
	case "hcl":
		err := hcl.Unmarshal(Sanitize(example.Data), dest)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown type %s %s", example.Type, example.Data)
	}

	return nil
}
