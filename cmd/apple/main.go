package main

import (
	"flag"
	"go_project_example/app_apple/apple_config"
	"go_project_example/init/apple_init"
	"log"
)

var (
	help       bool
	showConfig bool
)

func main() {
	if !apple_init.WithConfig() {
		return
	}

	flag.BoolVar(&help, "help", false, "show help")
	flag.BoolVar(&showConfig, "show", false, "show config")

	flag.Parse()

	if help {
		log.Println("I heard you need help")
	} else if showConfig {
		log.Printf("This is config:\n%+v", apple_config.GetAppConfig())
	} else {
		flag.Usage()
	}
}

/*
Use cmd:
	set env GO_PROJECT_EXAMPLE with `YOUR_DIR/go_project_example`
	cd YOUR_DIR/go_project_example/cmd/apple
	go build
	./apple
*/
