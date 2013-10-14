package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type CloudFormationOutput struct {
	Description string
}

type CloudFormationParameter struct {
	Default     string
	Description string
	Type        string
}

type CloudFormationResource struct {
	Type string
}

type CloudFormationTemplate struct {
	AWSTemplateFormatVersion string
	Description              string
	Mappings                 map[string]interface{}
	Outputs                  map[string]CloudFormationOutput
	Parameters               map[string]CloudFormationParameter
	Resources                map[string]CloudFormationResource
}

var acme bool
var show bool

func usage() {
	fmt.Fprintf(os.Stderr, "usage: cfn-summarize [-as] template\n")
	flag.PrintDefaults()
}

func main() {

	flag.Usage = usage

	// This is a feature that is probably only useful to @drocamor
	flag.BoolVar(&acme, "a", false, "Show acme addresses of items")
	flag.BoolVar(&show, "s", false, "Show the cfn-show command to detail this resource")

	flag.Parse()

	// You must provide at least one template file to summarize
	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	file, e := ioutil.ReadFile(flag.Arg(0))
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	var template CloudFormationTemplate
	json.Unmarshal(file, &template)

	fmt.Printf("Description: %s\n\n", template.Description)

	if template.Mappings != nil {
		fmt.Printf("Mappings:\n\n")
		for k, _ := range template.Mappings {
			fmt.Printf("%s\n", k)
			if acme {
				fmt.Printf(`%s:/%s\"[\ ]*:/`, flag.Arg(0), k)
				fmt.Printf("\n")
			}
			if show {
				fmt.Printf("cfn-show %s Mappings/%s\n", flag.Arg(0), k)
			}
			fmt.Printf("\n")
		}
	}

	if template.Outputs != nil {
		fmt.Printf("Outputs:\n\n")
		for k, v := range template.Outputs {
			fmt.Printf("%s\n", k)
			if acme {
				fmt.Printf(`%s:/%s\"[\ ]*:/`, flag.Arg(0), k)
				fmt.Printf("\n")
			}
			if show {
				fmt.Printf("cfn-show %s Outputs/%s\n", flag.Arg(0), k)
			}
			fmt.Printf("Description: %s\n", v.Description)
			fmt.Printf("\n")
		}
	}

	if template.Parameters != nil {
		fmt.Printf("Parameters:\n\n")
		for k, v := range template.Parameters {
			fmt.Printf("%s\n", k)
			if acme {
				fmt.Printf(`%s:/%s\"[\ ]*:/`, flag.Arg(0), k)
				fmt.Printf("\n")
			}
			if show {
				fmt.Printf("cfn-show %s Parameters/%s\n", flag.Arg(0), k)
			}
			fmt.Printf("Default: %s\n", v.Default)
			fmt.Printf("Description: %s\n", v.Description)
			fmt.Printf("Type: %s\n", v.Type)

			fmt.Printf("\n")
		}
	}

	if template.Resources != nil {
		fmt.Printf("Resources:\n\n")
		for k, v := range template.Resources {
			fmt.Printf("%s\n", k)
			if acme {
				fmt.Printf(`%s:/%s\"[\ ]*:/`, flag.Arg(0), k)
				fmt.Printf("\n")
			}
			if show {
				fmt.Printf("cfn-show %s Resources/%s\n", flag.Arg(0), k)
			}
			fmt.Printf("Type: %s\n", v.Type)
			fmt.Printf("\n")
		}
	}

}
