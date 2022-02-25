package db

import (
	"fmt"
	"go_dev/day17/demo/src/user-srv/entity"
	"time"
)


//insert into User
func InsertUser(Name string,Address string,Phone string) error {
	_ = time.Now().Format("2006-01-02")
	str :="insert into user(name,address,phone)value(?,?,?);"
	_, err := db.Exec(str, Name, Address, Phone)
	return err
}
//Delete 
func DeleteUser(id int32)(i int32,err error)  {
	str := "delete from user where id=?;"
	result, err := db.Exec(str, id)
	if err!=nil{
		fmt.Printf("错误%#v\n",err)
		return
	}
	_, err = result.RowsAffected()
	if err!=nil{
		fmt.Printf("错误%#v\n",err)
		return
	}
	i =id
	return
}
//Update
func UpdateUser(id int32,phone string)(i int32,err error)  {
	str :="update user set phone=? where id=?"
	_,err = db.Exec(str,phone,id)
	if err!=nil{
		fmt.Printf("错误%#v\n",err)
		return
	}
	i=id
	return
}
//Select
func SelectUser(id int32)(user *entity.User,err error){
	var u entity.User
	str :="select * from user where id=?;"
	err = db.QueryRow(str, id).Scan(&u.Id,&u.Name,&u.Address,&u.Phone)
	fmt.Printf("%#v\n",&u)
	if err!=nil{
		fmt.Printf("错误%#v\n",err)
	}
	user=&u
	return
}