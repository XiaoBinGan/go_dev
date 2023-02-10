package dao

import (
	"errors"
	"fmt"

	"gorm_connect/model"
	"gorm_connect/utils"
)

// add user
func AddUser(u *model.User) error {
	if err := utils.Db.Create(&u).Error; err != nil {
		fmt.Println("插入失败", err.Error())
		return err
	}
	return nil
}

// Checkuser
// if database has this user return nil,and if return error the database don't have this user
func Checkuser(u *model.User) error {
	var us model.User
	utils.Db.Where(&u).First(&us)
	if len(us.Username) != 0 {
		return errors.New("用户已存在")
	}
	return nil
}

// Userlogin
// if username and password are right return nil
// else return error
func Login(u *model.User) error {
	var us model.User
	utils.Db.Where(&u).First(&us)
	if u.Password != us.Password || u.Username != us.Username {
		return errors.New("用户名或者密码错误")
	}
	return nil
}

// userupdate
// if username and password are right return nil
// else return error
func ChangePassword(u *model.User, p string) error {
	// var us model.User
	err := utils.Db.Model(&u).Where("username=?", &u.Username).Update("password", p).Error
	if err != nil {
		return err
	}
	return nil
}
