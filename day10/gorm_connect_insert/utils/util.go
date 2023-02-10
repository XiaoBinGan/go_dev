package utils

import (
	"fmt"
	"time"

	// _ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"gorm_connect/conf"
)

var (
	// Db  *sql.DB
	Db    *gorm.DB
	DbErr error
)

func Init() { //int 强转string
	/****************************mysql s******************************************/
	// dns := conf.Cf.Username + ":" + conf.Cf.Password + "@tcp(" + conf.Cf.Host + ":" + strconv.Itoa(conf.Cf.Port) + ")/" + conf.Cf.Dbname
	// fmt.Println(dns)
	// Db, Err = sql.Open("mysql", dns)
	// if Err != nil {
	// 	panic(Err)
	// }
	// if Err = Db.Ping(); Err != nil {
	// 	panic(Err)
	// }
	// Db.SetMaxOpenConns(10)
	// Db.SetMaxIdleConns(2)
	// Db.SetConnMaxIdleTime(time.Second * 10)
	// fmt.Println("connection success")
	// defer Db.Close()
	/****************************mysql e******************************************/
	/****************************gormmysql s******************************************/
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", conf.Cf.Username, conf.Cf.Password, conf.Cf.Host, conf.Cf.Port, conf.Cf.Dbname)
	fmt.Println(dsn)

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic("连接数据库失败, error=" + DbErr.Error())
	}
	fmt.Println("database connect success")
	d := db.DB()
	d.SetMaxOpenConns(10)
	d.SetMaxIdleConns(2)
	d.SetConnMaxIdleTime(time.Second * 10)
	Db = db
	// defer Db.Close()
	// gorm维护了个连接池。注册单例之后就可以一直用。
}
