package db

import "testing"

func init(){
	//parseTime=true 将MySQL中的时间类型自动解析为go语言结构体中的时间类型
	dns :="root:@tcp(127.0.0.1:3306)/blogger?parseTime=true"
	err :=Init("mysql",dns)
	if err !=nil{
		panic(err)
	}
}

func TestGetCategoryById(t *testing.T) {
	category, err := GetCategoryById(1)
	if err!=nil{
		panic(err)
	}
	t.Logf("category:%#v",category)
}
func TestGetCategoryList(t *testing.T) {
	var categoryIds []int64
	categoryIds = append(categoryIds, 1, 2, 3, 5)
	list, err := GetCategoryList(categoryIds)
	if err!=nil{
		panic(err)
	}
	for _, v := range list {
		t.Logf("category:%#v\n",v)
	}
}
func TestGetAllcategoryList(t *testing.T) {
	list, err := GetAllcategoryList()
	if err!=nil{
		panic(err)
	}
	for _, v := range list {
		t.Logf("category:%#v\n",v)
	}
}