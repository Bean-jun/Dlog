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

type cert struct {
	Private    string `yaml:"private"`
	PrivateKey []byte
	Public     string `yaml:"public"`
	PublicKey  []byte
}

type server struct {
	SecretKey   string `yaml:"secretKey"`
	Port        string `yaml:"port"`
	TokenExpire int    `yaml:"token_expire"`
	Cert        cert   `yaml:"cert"`
}

type Config struct {
	Mysql    mysql  `yaml:"mysql"`
	Server   server `yaml:"server"`
	Debug    bool   `yaml:"debug"`
	Version  int
	RootPath string
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

	Conf.Version, err = ParseVersion()
	if err != nil {
		panic(err.Error())
	}

	// set root path
	dir, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	Conf.RootPath = dir

	// read cert
	privKey, err := os.ReadFile(Conf.Server.Cert.Private)
	if err != nil {
		panic(err.Error())
	}
	Conf.Server.Cert.PrivateKey = privKey
	pubKey, err := os.ReadFile(Conf.Server.Cert.Public)
	if err != nil {
		panic(err.Error())
	}
	Conf.Server.Cert.PublicKey = pubKey
}
