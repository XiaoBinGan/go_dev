package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //引入进来但是不直接操作
	"github.com/jmoiron/sqlx"           //使用这个库来对数据库进行操作
)

var (
	DB *sqlx.DB
)

/**
@params{
	tp   connect type
	dns  root:admin:@tcp(127.0.0.1:3306)/testdatebese
}
explain:
	1.create sql connect
	2.return error
	3.ping db is connect？
	4.return error
	5.the max open thread
	6.the max free thread
	7.return nil
*/

func Init(tp,dns string) error  {
	var err error
	DB, err = sqlx.Open(tp, dns)
	if err!=nil{
		return err
	}
	fmt.Println(DB)
	err = DB.Ping()
	if err!=nil{
		return err
	}
	DB.SetMaxOpenConns(1000)
	DB.SetMaxIdleConns(20)
	return nil
}



