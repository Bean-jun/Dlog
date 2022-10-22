package main

import (
	"github.com/Bean-jun/Dlog/dao"
	"log"

	"github.com/Bean-jun/Dlog/pkg"
	"github.com/Bean-jun/Dlog/router"
)

func InitEnv() {
	pkg.InitConfig("conf.yaml")
	dao.InitDB()
}

func main() {
	InitEnv()
	route := router.InitRouter()
	log.Fatal(route.Run(pkg.Conf.Server.Port))
}
