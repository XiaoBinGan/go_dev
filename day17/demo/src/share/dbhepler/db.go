package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go_dev/day17/demo/src/share/config"
	"log"
	"math/rand"
	"time"
)

//建表语句
//因为是ID ，所以是自增无符号的 Int
var schema = `
CREATE TABLE IF NOT EXISTS user (
	id INT UNSIGNED AUTO_INCREMENT,
	name VARCHAR(20),
	address VARCHAR(20),
	phone VARCHAR(15),
	PRIMARY KEY(id)
)
`

//对应表的结构体
type User struct {
	Id   		int32  `db:"id"`
	Name 		string `db:"name"`
	Address 	string `db:"address"`
	Phone 		string `db:"phone"`
}

func main()  {
	//打开数据库。返回错误信息
	//mysql是驱动名称,上面需要侵入数据库驱动包
	//第二参数是指定连接到哪个数据库，在share/config中配置
	db, err := sqlx.Connect("mysql", config.MysqlDSN)
	if err!=nil{
		log.Fatalln(err)
	}
	//执行建表语句
	db.MustExec(schema)
	//开启事务
	tx := db.MustBegin()
	//设置随机数种子，可以保证每次随机都是随机的
	rand.Seed(time.Now().UnixNano())
	//事务执行SQL插入
	//创建GetRandomString()方法，按指定个数生成随机字符
	tx.MustExec(`INSERT INTO user (id,name,address,phone) VALUES (?,?,?,?)`,
		nil, GetRandomString(10), "beijing"+GetRandomString(10),
		"1591"+GetRandomString(7))
	//提交事务
	err = tx.Commit()
	if err!=nil{
		//回滚
		_=tx.Rollback()
	}
	fmt.Println("执行完毕！")
}
func GetRandomString(leng int)string {
	str :="123145646489778123wqdaerq3ewqeqw213e12q"
	bytes :=[]byte(str)
	result :=[]byte{}
	//用系统时间生成随机数
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//指定个数添加到返回结果中
	for i := 0; i <leng ; i++ {
		result=append(result,bytes[r.Intn(len(bytes))])
	}
	return string(result)
}


















