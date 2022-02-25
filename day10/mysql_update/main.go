package main

import (
	"fmt"
	 _"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)
type Person struct {
	UserId   int     `db:"user_id"`
	Username string	 `db:"username"`
	Sex		 string  `db:"sex"`
	Email    string  `db:"email"`
}

var Db *sqlx.DB


func init()  {
	database,err :=sqlx.Open("mysql","root:@tcp(127.0.0.1:3306)/test")
	if err!=nil{
		fmt.Println("open mysql failed,err:",err)
		return
	}
	fmt.Println(database)
	Db = database
}

func main() {
	_,err :=Db.Exec("update person set username=? where user_id=?","str00000001",1)
	if err!=nil{
		fmt.Println("")
		return
	}
		
}