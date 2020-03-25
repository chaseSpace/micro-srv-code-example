package banana_init

import (
	"github.com/BurntSushi/toml"
	"go_project_example/app_apple/apple_config"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var (
	webConfigFileName = "app_apple.toml"
)

func WithConfig() (ok bool) {
	// get project major directory
	projectDir := os.Getenv("GO_PROJECT_EXAMPLE")
	if projectDir == "" {
		log.Println("you must set ENV `GO_PROJECT_EXAMPLE` with you project directory")
		return
	}
	// put them together to get path of config file
	confPath := filepath.Join(projectDir, "configs", webConfigFileName)

	var conf apple_config.Config
	confData, _ := ioutil.ReadFile(confPath)

	if _, err := toml.Decode(string(confData), &conf); err != nil {
		log.Panicf("toml parse config err:%v", err)
	}
	// set global config
	apple_config.InitConfig(&conf)
	return true
}
