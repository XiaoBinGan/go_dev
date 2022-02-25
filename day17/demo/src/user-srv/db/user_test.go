package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //引入做驱动
	"go_dev/day17/demo/src/share/config"
	"testing"
)

func init()  {
	Init( config.MysqlDSN)
}


func TestInsertUser(t *testing.T) {
	err := InsertUser("Alben1", "2342", "hotamail222")
	if err!=nil{
		t.Logf("%#v",err)
	}
}

func TestDeleteUser(t *testing.T) {
	i, err := DeleteUser(4)
	if err!=nil{
		t.Logf("%#v",err)
	}
	fmt.Println(i)
}
func TestUpdateUser(t *testing.T) {
	i, err := UpdateUser(5,"741852963")
	if err!=nil{
		t.Logf("%#v",err)
	}
	fmt.Println(i)
}

func TestSelectUser(t *testing.T) {
	selectUser, err := SelectUser(2)
	if err!=nil{
		t.Logf("#{err}")
	}
	fmt.Print(selectUser)
}