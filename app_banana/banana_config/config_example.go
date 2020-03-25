package banana_config

import (
	"log"
	"time"
)

type Config struct {
	Age        int
	Cats       []string
	Pi         float64
	Perfection []int
	DOB        time.Time
}

var conf *Config

func GetAppConfig() *Config {
	if conf == nil {
		log.Panic("not init yet.")
	}
	return conf
}

func InitConfig(c *Config) {
	conf = c
}
