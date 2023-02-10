package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"gorm_connect/dao"
	"gorm_connect/model"
	"gorm_connect/utils"
)

// func TestAddUser(t *testing.T) {
// 	conf.Init()  //init config
// 	utils.Init() //init database
// 	fmt.Println("添加用户")
// 	var us = model.User{}
// 	var u = model.User{
// 		Username: "认为二位二位",
// 		Password: "12312",
// 	}
// 	if err := service.AddUser(u.Username, u.Password); err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Printf("where:%#v", us)
// }

func TestPost(t *testing.T) {
	// type user struct {
	// 	username string
	// 	password string
	// }
	url := "http://127.0.0.1:8080/v1/adduser"
	// 表单数据
	// contentType := "application/x-www-form-urlencoded"
	//data := "name=小王子&age=18"
	// json
	contentType := "application/json"
	// contentType := "application/form-data"
	data := `{"username":"二位","password":"eqweqweqw"}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))

}

func TestJTW(t *testing.T) {
	s, err := utils.GenToken("name")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("token", s)

	mc, err := utils.ParseToken(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("tokenparse%#v", mc)

	// var mySigningKey = []byte("supremind.com")
	// // 创建 Claims
	// claims := &jwt.RegisteredClaims{
	// 	ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 过期时间
	// 	Issuer:    "wujiahao",                                             // 签发人
	// }
	// // 生成token对象
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// // 生成签名字符串
	// s, err := token.SignedString(mySigningKey)
	// if err == nil {
	// 	fmt.Println("s:", s)
	// }
}

func TestCahngePassword(t *testing.T) {
	var us = &model.User{
		Username: "qwewqeqwe",
		Password: "123123123",
	}
	err := dao.ChangePassword(us, "改好了")
	if err != nil {
		fmt.Println("err:", err.Error())
	}
}
