package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("mysql open failed ,err:", err)
		return
	}
	Db = database
}

func main() {
	_, err := Db.Exec("delete from person where user_id=?", 2)
	if err != nil {
		fmt.Println("exec failed err:", err)
		return
	}
	fmt.Println("delete succ")
}
