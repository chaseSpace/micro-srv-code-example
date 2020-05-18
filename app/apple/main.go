package main

import (
	"go_project_template/app/apple/internal"
	"go_project_template/config/apple_conf"
	"go_project_template/init/apple_init"
	"log"
	"time"
)

// customize your env
const GoProjectDir = "GO_PROJECT_DIR"

func main() {
	apple_init.MustInit(GoProjectDir)

	conf := apple_conf.GetAppConfig()
	log.Printf("%+v", *conf)

	log.Println(internal.AppleHi())

	log.Println("apple is running~")

	time.Sleep(2 * time.Second)

	log.Println("apple exit~")
}
