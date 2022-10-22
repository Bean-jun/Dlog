package pkg

import (
	"os"

	"gopkg.in/yaml.v3"
)

type mysql struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
}

type server struct {
	Port string `yaml:"port"`
}

type Config struct {
	Mysql  mysql  `yaml:"mysql"`
	Server server `yaml:"server"`
	Debug  bool   `yaml:"debug"`
}

var (
	Conf = &Config{}
)

func InitConfig(path string) {
	confData, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}

	err = yaml.Unmarshal(confData, Conf)
	if err != nil {
		panic(err.Error())
	}
}
