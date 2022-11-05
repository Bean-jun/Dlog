package main

import (
	"fmt"
	"log"

	"github.com/Bean-jun/Dlog/dao"
	"github.com/Bean-jun/Dlog/router"

	"github.com/Bean-jun/Dlog/pkg"
)

func InitEnv() {
	pkg.InitConfig("conf.yaml")
	dao.InitDB()
}

func welcome() {
	fmt.Println("\t\t--------------------------------")
	fmt.Printf("\t\t|hello Dlog V%s|\n", pkg.Version)
	fmt.Println("\t\t--------------------------------")
}

// @title Dlog
// @version 1.0
// @description Dlog后端api服务
// @contact.name Bean-jun
// @contact.url https://github.com/Bean-jun
// @contact.email 1342104001@qq.com
// @license.name MIT license
// @host 192.168.2.100:9090
// @BasePath /
func main() {
	InitEnv()
	route := router.InitRouter()
	welcome()
	log.Fatal(route.Run(pkg.Conf.Server.Port))
}
