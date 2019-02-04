package main

import (
	"fmt"
	"github.com/bunniesandbeatings/symphonic"
	"github.com/jessevdk/go-flags"
	"os"
)

var config symphonic.Config

var parser = flags.NewParser(&config, flags.Default)

func main() {
	fmt.Println("Symphonic 12 Factor Microservices Orchestrator")

	var config symphonic.Config
	_, err := parser.Parse()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	if config.Debug {
		fmt.Printf("Config: %+v", config)
	}

}


