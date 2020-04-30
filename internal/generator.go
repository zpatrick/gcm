package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

type TemplateContext struct {
	Imports   []string
	Providers []ProviderContext
}

type ProviderContext struct {
	Type            string
	TypeNamePrefix  string
	Variable        string
	ZeroValue       string
	ValidateInSet   bool
	ValidateBetween bool
}

type ValidatorContext struct {
	Enabled bool
}

func main() {
	path := flag.String("template", "provider.go.template", "path to template file")
	flag.Parse()

	content, err := ioutil.ReadFile(*path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	t, err := template.New("").Parse(string(content))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tc := TemplateContext{
		Imports: []string{"time"},
		Providers: []ProviderContext{
			{
				Type:           "interface{}",
				TypeNamePrefix: "",
				Variable:       "v",
				ZeroValue:      "nil",
				ValidateInSet:  true,
			},
			{
				Type:           "bool",
				TypeNamePrefix: "Bool",
				Variable:       "b",
				ZeroValue:      "false",
			},
			{
				Type:            "int",
				TypeNamePrefix:  "Int",
				Variable:        "i",
				ZeroValue:       "0",
				ValidateInSet:   true,
				ValidateBetween: true,
			},
			{
				Type:            "int64",
				TypeNamePrefix:  "Int64",
				Variable:        "i",
				ZeroValue:       "0",
				ValidateInSet:   true,
				ValidateBetween: true,
			},
			{
				Type:            "float64",
				TypeNamePrefix:  "Float64",
				Variable:        "f",
				ZeroValue:       "0",
				ValidateInSet:   true,
				ValidateBetween: true,
			},
			{
				Type:           "string",
				TypeNamePrefix: "String",
				Variable:       "s",
				ZeroValue:      "\"\"",
				ValidateInSet:  true,
			},
			{
				Type:            "time.Duration",
				TypeNamePrefix:  "Duration",
				Variable:        "d",
				ZeroValue:       "0",
				ValidateInSet:   true,
				ValidateBetween: true,
			},
		},
	}

	if err := t.Execute(os.Stdout, tc); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
