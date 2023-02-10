package conf

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

// dbconfig
type DbConf struct {
	Username string `ini:"username"` //账号
	Password string `ini:"password"` //密码
	Host     string `ini:"host"`     //数据库地址，可以是Ip或者域名
	Port     int    `ini:"port"`     //数据库端口
	Dbname   string `ini:"dbname"`   //数据库名
	Timeout  int    `ini:"timeout"`  //超时时间
}

var (
	Cf  *DbConf
	Err error
)

// ini mysql database from ini.conf
func Init() {
	f, err := ini.Load("./conf/db.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	Cf = new(DbConf)
	if Err = f.Section("DB").MapTo(Cf); err != nil {
		os.Exit(1)
	}
	fmt.Printf("DbConf ini:%#v", Cf)
	fmt.Println("Usernam:", f.Section("DB").Key("username").String())

}
