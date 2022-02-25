package db

import (
	_ "github.com/go-sql-driver/mysql" //引入做驱动
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)
//init mysql connect
func Init(mysqlDSN string)  {
	db = sqlx.MustConnect("mysql", mysqlDSN)
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(3)
}