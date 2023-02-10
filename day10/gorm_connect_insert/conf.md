 ### 1

 ### 1、创建conf.ini文件

```ini
app_mode = development ;//环境
[DB]  ;//selection
	; //配置MySQL连接参数
username = root  ;//账号
password = ;//密码
host = 127.0.0.1 ;//数据库地址，可以是Ip或者域名
port = 3306        ;//数据库端口
dbname = gormtest  ;//数据库名

```



### 2、读取ini文件，读取配置

```go
// ini mysql database from ini.conf
func init() {
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

```



### 3、初始化db

```go

package utils

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"gorm_connect/conf"
)

var (
	Db  *sql.DB
	Err error
)

func init() { //int 强转string
	dns := conf.Cf.Username + "@" + conf.Cf.Password + "tcp(" + conf.Cf.Host + ":" + strconv.Itoa(conf.Cf.Port) + ")/" + conf.Cf.Dbname
	fmt.Println(dns)
	Db, Err = sql.Open("mysql", dns)
	if Err != nil {
		panic(Err)
	}
	if Err = Db.Ping(); Err != nil {
		panic(Err)
	}
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(2)
	Db.SetConnMaxIdleTime(time.Second * 10)
	fmt.Println("connection success")
	defer Db.Close()
}

```

### 4、创建dao层

