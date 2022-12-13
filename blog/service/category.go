package service

import (
	"fmt"
	"blog/dao/db"
	"blog/model"
)
/**
the GetAKKCateGategoryList service is get all category for categoryListitem
*/
func GetAllCateGategoryList() (categoryList []*model.Categroy,err error) {
	categoryList, err = db.GetAllcategoryList()
	fmt.Printf("err:%#v",err)
	fmt.Println("categoryList:",categoryList)
	if err!=nil{
		return
	}
	return
}