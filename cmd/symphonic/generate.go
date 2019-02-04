package main

import (
	"encoding/json"
	"fmt"
	"github.com/bunniesandbeatings/symphonic/service"
	"os"
)

var generateOpts Generate

type Generate struct {
}

func (cmd *Generate) Execute(args []string) error {

	serviceDef := service.Service{
		Name:"auth",
		Route: service.Route{
			Path: "auth",
		},
		Clean: []string{"make", "clean"},
		Deps: []string{"make", "deps"},


	}

	serviceDefJSON, err := json.Marshal(serviceDef)

	if err != nil {
		fmt.Println("Could not marshall JSON")
		os.Exit(1)
	}

	fmt.Println(string(serviceDefJSON))
	return nil
}

func init() {
	_, err := parser.AddCommand("generate", "Generate samples", "Generate config samples", &generateOpts)
	if err != nil {
		fmt.Println("Could not add generate command")
		os.Exit(1)
	}
}
