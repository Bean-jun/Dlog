package dao

import (
	"fmt"

	"github.com/Bean-jun/Dlog/dao/entity"
	"github.com/Bean-jun/Dlog/pkg"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB         *gorm.DB
	RDB        *redis.Client
	err        error
	entityList = []interface{}{&entity.UserEntity{}}
)

func InitDB() {
	InitMySQL()
	InitRedis()
}

func InitMySQL() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		pkg.Conf.DB.Mysql.Username,
		pkg.Conf.DB.Mysql.Password,
		pkg.Conf.DB.Mysql.Host,
		pkg.Conf.DB.Mysql.Port,
		pkg.Conf.DB.Mysql.Dbname,
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

func InitRedis() {
	dsn := fmt.Sprintf("%s:%d", pkg.Conf.DB.Redis.Host, pkg.Conf.DB.Redis.Port)
	RDB = redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: pkg.Conf.DB.Redis.Password,
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})

	_, err = RDB.Ping().Result()
	if err != nil {
		panic(err)
	}
}
