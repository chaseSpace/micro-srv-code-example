package apple_init

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"go_project_template/config/apple_conf"
	"go_project_template/pkg"
	"go_project_template/util"
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

	flag.StringVar(&cmdConfPath, "conf", "", "string%sconfig path")
	flag.StringVar(&cmdConfPath, "c", "", "string%sconfig path")

	flag.BoolVar(&cmdHelp, "help", false, "bool%sshow help")
	flag.BoolVar(&cmdHelp, "h", false, "bool%sshow help")

	flag.Parse()

	if cmdHelp || len(os.Args) == 1 {
		split := "   "
		flag.Usage = func() {
			_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
			complete := "" // complete flag, eg -conf / -help
			single := ""   // single char flag, eg -c / -h
			flag.VisitAll(func(i *flag.Flag) {
				if len(i.Name) > 1 {
					complete += fmt.Sprintf("%s-%s   %s\n", split, i.Name, fmt.Sprintf(i.Usage, split))
				} else {
					single += fmt.Sprintf("%s-%s   %s\n", split, i.Name, fmt.Sprintf(i.Usage, split))
				}
			})
			_, _ = fmt.Fprintf(flag.CommandLine.Output(), "%sSimplify flag:\n%s\n", complete, single)
		}
		flag.Usage()
		/*
			$ go run .
			Usage of /path/to/NAME.exe:
			   -conf   string   config path
			   -help   bool   show help
			simplify flag:
			   -c   string   config path
			   -h   bool   show help
		*/
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
			"---Input `go run . help` to see cmd option\n"+
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
