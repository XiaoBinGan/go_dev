package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB
var err error
//`id`          integer       NOT NULL PRIMARY KEY AUTOINCREMENT,
//`version`     varchar(50)  DEFAULT NULL,
//`description` varchar(200) DEFAULT NULL,
//`script`      varchar(1000) NOT NULL,
//`create_time` datetime,
//`update_time` datetime
type D struct {
	id			int
	version     string
	description string
	script      string
	create_time time.Time
	update_time time.Time


}


// TableName setting the table name
func (D) TableName() string {
	return "schema_history"
}


func main()  {
	dbUrl := "/Users/wujiahao/Desktop/105.db"
	db, err := gorm.Open(sqlite.Open(dbUrl), &gorm.Config{})
	if err!=nil{
		log.Fatalf("failed to connect database")
	}
	var demo D
	rows,err :=db.Model(&D{}).Rows()
	defer  rows.Close()
	if err!=nil{
		panic(err)
	}
	for rows.Next(){
		db.ScanRows(rows,&demo)
		fmt.Println(demo)
	}
}