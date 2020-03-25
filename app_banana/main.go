package main

import (
	"go_project_example/app_banana/banana_config"
	"go_project_example/init/banana_init"
	"log"
)

func main() {
	banana_init.WithConfig()

	conf := banana_config.GetAppConfig()
	log.Printf("%+v", *conf)

	log.Println("banana is running~")
}
