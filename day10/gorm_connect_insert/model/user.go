package model

// User users information
/**
这里申明一个结构体
form tag    form表单提交
json tag	post提交json数据格式
binding:"required"  必传字段
分别用来告诉 context.ShouldBind()什么请求类型的参数对应结构体的什么名称
ShouldBind()方法使用反射来对应结构体内部参数和请求参数，所以请注意结构体内部的参数尽量大写开头
ShouldBind可以对应任何请求类型的参数只要在结构体的参数后面加上相对于的tag
*/
type User struct {
	ID         int64  //key
	Username   string `gorm:"column:username" form:"username" json:"username" binding:"required"`
	Password   string `gorm:"column:password" form:"password" json:"password" binding:"required"`
	CreateTime int64  `gorm:"column:createtime"`
}

// table name设置表名，可以通过给struct类型定义 TableName函数，返回当前struct绑定的mysql表名是什么
func (u *User) TableName() string {
	//绑定MYSQL表名为users
	return "users"
}
