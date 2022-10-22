package utils

import (
	"fmt"
	"strings"

	"github.com/Bean-jun/Dlog/pkg"
	"github.com/dgrijalva/jwt-go"
)

// GenerateToken 生成 token
func GenerateToken(data jwt.MapClaims) (token string, err error) {
	// 创建一个新的令牌对象，指定签名方法和声明
	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	// 使用密码签名并获得完整的编码令牌作为字符串
	token, err = tokenString.SignedString([]byte(pkg.Conf.Server.SecretKey))
	return
}

func splitToken(tokenstr string) string {
	return strings.Split(tokenstr, "Bearer ")[1]
}

// ParseToken 解析 token
func ParseToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(splitToken(tokenStr), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("异常签名: %v", token.Header["alg"])
		}
		return []byte(pkg.Conf.Server.SecretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
