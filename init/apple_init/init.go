package apple_init

import (
	"flag"
	"github.com/BurntSushi/toml"
	"go_project_example/config/apple_conf"
	"go_project_example/pkg"
	"go_project_example/util"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	projectDirEnvKey string
	projectDirEnv    string

	// config的后面部分路径，与环境变量组合使用， file = $projDirEnvKey/filepath.join(confLaterPath)
	confLaterPath = []string{"staticfile", "config", "apple.toml"}
	//assetLaterPath     = []string{"staticfile", "asset"}
)

var (
	cmdConfPath string
	cmdHelp     bool
)

func init() {
	flag.StringVar(&cmdConfPath, "conf", "", "config path")
	flag.BoolVar(&cmdHelp, "cmdHelp", false, "show cmdHelp")

	flag.Parse()

	if cmdHelp {
		flag.Usage()
		os.Exit(1)
	}
}

func mustReadFile(path string) []byte {
	confData, err := ioutil.ReadFile(path)
	util.PanicIfErr(err)
	return confData
}

func getConfFile() []byte {
	// from cmd
	if cmdConfPath != "" {
		absPath, err := filepath.Abs(cmdConfPath)
		util.PanicIfErr(err)
		return mustReadFile(absPath)
	}

	if projectDirEnv == "" {
		pkg.EnvNotSetErr.Panic("Please provide config file path by <cmd/environment>\n"+
			"---Input `go run . cmdHelp` to see cmd option\n"+
			"---Set env `%s` to solve most of init problems", projectDirEnvKey)
	}
	// from env
	confPath := filepath.Join(append([]string{projectDirEnv}, confLaterPath...)...)
	return mustReadFile(confPath)
}

func MustInitConfig() {
	var conf apple_conf.Config
	_, err := toml.Decode(string(getConfFile()), &conf)
	util.PanicIfErr(err)

	// set global config
	apple_conf.InitConfig(&conf)
}

func MustInit(projDirEnvKey string) {
	projectDirEnvKey = projDirEnvKey
	projectDirEnv = os.Getenv(projDirEnvKey)

	pkg.MustInit(projDirEnvKey, projectDirEnv)
	MustInitConfig()
}
