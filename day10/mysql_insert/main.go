package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"  //引入进来但是不直接操作
	"github.com/jmoiron/sqlx"			//使用这个库来对数据库进行操作
)
type Person struct {
	UserId   int    `db:"user_id"`    //db代表的是最终存储到数据库的字段 类似json.Marshal json tag将对应的字段转成映射的tag
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}


var Db *sqlx.DB                       //实力化一个可以操作数据库对象不需要用户手动再去新建连接池

func init() {			  //数据库类型	  账户:密码@tcp协议（ip地址:端口号）/表名称									   //初始化数据库连接		
	database, err := sqlx.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")  //创建数据库连接
	if err != nil {															
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database														   //将创建好的连接出入到全局
}

func main() {
					  //插入	到   什么表	字段1	字段2    字段3   占位符对应相对字段  字段1值     字段2值     字段3值
	r, err := Db.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com") //执行sql语句
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()//返回最后插入的记录的ID
	if err != nil { 
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("insert succ:", id)
}
// kafka-topics --create --zookeeper localhost:2181 --partitions 1 --topic test