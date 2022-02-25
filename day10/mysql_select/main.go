 package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" //引入进来但是不直接操作
	"github.com/jmoiron/sqlx"          //使用这个库来对数据库进行操作
)

type Person struct {
	UserId   int    `db:"user_id"` //db代表的是最终存储到数据库的字段 类似json.Marshal json tag将对应的字段转成映射的tag
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

var DB *sqlx.DB

func init() { //init函数会在main函数执行前执行
	database, err := sqlx.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	DB = database
}
func main() {
	var person []Person
	err := DB.Select(&person, "select * from person where user_id=?", 1)
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	fmt.Println("select succ:", person)

}
