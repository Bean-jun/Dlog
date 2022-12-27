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

type redis struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

type db struct {
	Mysql mysql `yaml:"mysql"`
	Redis redis `yaml:"redis"`
}

type cert struct {
	Private    string `yaml:"private"`
	PrivateKey []byte
	Public     string `yaml:"public"`
	PublicKey  []byte
}

type server struct {
	SecretKey string `yaml:"secretKey"`
	Port      string `yaml:"port"`
	Cert      cert   `yaml:"cert"`
}

type account struct {
	TokenExpire                     int `yaml:"token_expire"`
	VerifyCodeExpire                int `yaml:"verify_code_expire"`
	LoginErrNum                     int `yaml:"login_err_num"`
	LoginErrTips                    int `yaml:"login_err_tips"`
	LoginErrInterval                int `yaml:"login_err_interval"`
	LoginErrLock                    int `yaml:"login_err_lock"`
	AcountMaxNotActive              int `yaml:"acount_max_not_active"`
	AcountMaxModifyPasswordInterval int `yaml:"acount_max_modify_password_interval"`
}

type Config struct {
	Server   server  `yaml:"server"`
	Account  account `yaml:"account"`
	DB       db      `yaml:"db"`
	Debug    bool    `yaml:"debug"`
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
