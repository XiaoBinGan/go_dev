package repository

import (
	"github.com/jinzhu/gorm"
	"user/domain/model"
)

type IUserRepository interface {
	//init dataList
	InitTable()error
	//query user info by user_name
	FindUserByName(string)(*model.User,error)
	//query user info by ID
	FindUserByID(int64)(*model.User,error)
	//create user
	CreateUser(*model.User)(int64,error)
	//delete user by ID
	DeleteUserByID(int64)error
	//update user info
	UpdateUser(*model.User)error
	//find all
	FindAll()([]model.User,error)
}
//create UserRepository
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlDb:db}
}

type UserRepository struct {
	 mysqlDb *gorm.DB
}

//init list
func (u *UserRepository)InitTable()error {
	return u.mysqlDb.CreateTable(&model.User{}).Error
}
//Finde user by name
func (u *UserRepository)FindUserByName(name string)(user *model.User,err error) {
	user =&model.User{}
	return user,u.mysqlDb.Where("user_name=?",name).Find(user).Error
}
//Finde user by Id
func (u *UserRepository)FindUserByID(userId int64)(user *model.User,err error){
	user = &model.User{}
	return user,u.mysqlDb.First(user,userId).Error
}
//create user
func (u *UserRepository)CreateUser(user *model.User)(userId int64,err error){
	return user.ID,u.mysqlDb.Create(user).Error
}
//delete user by Id
func (u *UserRepository)DeleteUserByID(userId int64)error {
	return u.mysqlDb.Where("id=?",userId).Delete(&model.User{}).Error
}
//update user info
func(u *UserRepository)UpdateUser(user *model.User)error{
	return u.mysqlDb.Model(user).Update(&user).Error
}
//FindAll
func (u *UserRepository)FindAll()(userAll []model.User,err error) {
	return userAll,u.mysqlDb.Find(&userAll).Error
}
