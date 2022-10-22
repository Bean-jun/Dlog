package dao

import (
	"fmt"

	"github.com/Bean-jun/Dlog/dao/entity"
	"github.com/Bean-jun/Dlog/pkg"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB         *gorm.DB
	err        error
	entityList = []interface{}{&entity.UserEntity{}}
)

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		pkg.Conf.Mysql.Username,
		pkg.Conf.Mysql.Password,
		pkg.Conf.Mysql.Host,
		pkg.Conf.Mysql.Port,
		pkg.Conf.Mysql.Dbname,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	err := DB.AutoMigrate(entityList...)
	if err != nil {
		panic(err)
	}
}
