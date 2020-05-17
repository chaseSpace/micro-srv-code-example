package main

import (
	"go_project_example/init/apple_init"
	"log"
)

// customize your env
const GoProjectDir = "GO_PROJECT_DIR"

func main() {
	/*
		cmd可以是app的运行时工具，也可以app的辅助工具
	*/
	apple_init.MustInit(GoProjectDir)
	log.Printf("cmd run success!")
}

/*
Use cmd:
	set env GO_PROJECT_EXAMPLE with `/path/to/this-project`
	cd /path/to/go_project_example/cmd/apple
	go build -o applectl
	./applectl help
*/
