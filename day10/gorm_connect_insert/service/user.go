package service

import (
	"errors"
	"time"

	"gorm_connect/dao"
	"gorm_connect/model"
)

/*
*
@Name Adduser
@param username string
@param password string
@return error    error
*/
func AddUser(username string, password string) error {
	tmpU := model.User{
		Username: username,
	}
	u := model.User{
		Username:   username,
		Password:   password,
		CreateTime: time.Now().Unix(),
	}
	if err := dao.Checkuser(&tmpU); err != nil {
		return errors.New("用户已存在")
	}
	if err := dao.AddUser(&u); err != nil {
		return err
	}
	return nil
}

/*
*
@Name Login
@param username string
@param password string
@return error    error
@detail
*/
func Login(username string, password string) error {
	tmpU := model.User{
		Username: username,
		Password: password,
	}
	if err := dao.Login(&tmpU); err != nil {
		return err
	}
	return nil
}
/*
*
@Name ChangePassword
@param username string
@param password string
@param ChangePassword string
@return error    error
@detail
*/