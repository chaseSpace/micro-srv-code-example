package main

// go1.13中模块路径第一个位置必需符合域名语法，之前的版本不受影响，所以不能使用`-`
import (
	"go_project_example/app_apple/apple_config"
	"go_project_example/init/apple_init"
	"log"
)

func main() {
	if !apple_init.WithConfig() {
		return
	}

	conf := apple_config.GetAppConfig()
	log.Printf("%+v", *conf)

	log.Println("apple is running~")
}
