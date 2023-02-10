package utils

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//1 生成token
//2 解析token
//3 过滤不需要鉴权的接口
//4 设置过期时间

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	Username string `json:"name"`
	jwt.StandardClaims
}

// JWT过期时间
const TokenExpireDuration = time.Hour * 2

var MySecret = []byte("supremind.com")

// 生成token
func GenToken(username string) (string, error) {
	c := MyClaims{
		username, //自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), //过期时间
			Issuer:    "supremind",
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	signedString, err := token.SignedString(MySecret)
	if err != nil {
		return "生成失败", nil
	}
	return signedString, nil
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	//解析username
	// 校验token
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
