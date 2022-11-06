package utils

import (
	"bytes"
	"encoding/base64"
	"image/jpeg"
	"log"
	"time"

	"github.com/Bean-jun/Dlog/pkg"
	"github.com/go-redis/redis"
	uuid "github.com/gofrs/uuid"
	"github.com/vcqr/captcha"
)

func GenerateCaptcha(RDB *redis.Client) (string, string) {
	cp := captcha.NewCaptcha(120, 40, 4)
	cp.SetFontPath("utils/") //指定字体目录
	cp.SetFontName("ubuntu") //指定字体名字
	cp.SetMode(1)            //1：设置为简单的数学算术运算公式； 其他为普通字符串
	code, img := cp.OutPut()

	// 获取UUID
	uuidStr := uuid.Must(uuid.NewV4()).String()

	log.Println("[INFO]: ", uuidStr, code)
	emptyBuff := bytes.NewBuffer(nil)       //开辟一个新的空buff
	err := jpeg.Encode(emptyBuff, img, nil) //img写入到buff
	if err != nil {
		panic(err)
	}
	dist := make([]byte, 50000)                        //开辟存储空间
	base64.StdEncoding.Encode(dist, emptyBuff.Bytes()) //buff转成base64
	index := bytes.IndexByte(dist, 0)
	dist_u := dist[0:index]
	data := string(dist_u) //输出图片base64(type = []byte)

	// 写入redis
	err = RDB.Set(uuidStr, code, time.Duration(pkg.Conf.Account.VerifyCodeExpire)*time.Second).Err()

	// 返回结果
	if err != nil {
		panic(err)
	}
	return uuidStr, data
}

// 校验验证码
func VerifyCode(RDB *redis.Client, uuidStr string, code string) bool {
	if val, err := RDB.Get(uuidStr).Result(); err == nil {
		if val == code {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
