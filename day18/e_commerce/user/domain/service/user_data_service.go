package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"user/domain/model"
	"user/domain/repository"
)

type IUserDataService interface {
	AddUser(*model.User)(int64,error)
	DeleteUser(int64) error
	UpdateUser(user *model.User,isChangePwd bool)(err error)
	FindUserByName(string)(*model.User,error)
	CheckPwd(userName string,pwd string)(isOk bool,err error)
}
//创建实例
func NewUserDataService(userRepository repository.IUserRepository) IUserDataService {
	return &UserDataService{UserRepository:userRepository}
}
// user Ddata manager
type UserDataService struct {
	UserRepository repository.IUserRepository
}
//encryption password
func GeneratePassword(userPassword string)([]byte,error){
	return bcrypt.GenerateFromPassword([]byte(userPassword),bcrypt.DefaultCost)
}
//validate Password 
func ValidatePassword(userPassword string,hashed string)(isOk bool,err error)  {
	if err= bcrypt.CompareHashAndPassword([]byte(hashed),[]byte(userPassword));err!=nil{
		return false, errors.New("密码对比错误")
	}
	return true,nil
}
//insert user
func (u *UserDataService)AddUser(user *model.User)(userID  int64,err error)  {
	pwdByte, err := GeneratePassword(user.HashPassword)
	if err!=nil{
		return user.ID,err
	}
	user.HashPassword=string(pwdByte)
	return u.UserRepository.CreateUser(user)
}
//delete user
func (u *UserDataService)DeleteUser(userId int64)error{
	return u.UserRepository.DeleteUserByID(userId)
}
//update user
func (u *UserDataService)UpdateUser(user *model.User,isChangePwd bool)(err error){
	//is Change password
	if isChangePwd{
		//encryption the password
		bytepwd, err := GeneratePassword(user.HashPassword)
		if err!=nil{
			return err
		}
		user.HashPassword=string(bytepwd)
	}
	return u.UserRepository.UpdateUser(user)
}
func (u *UserDataService)FindUserByName(userName string)(user *model.User,err error){
	return u.UserRepository.FindUserByName(userName)
}

//CheckPwd(userName string,pwd string)
func (u *UserDataService)CheckPwd(userName string,pwd string)(isOk bool,err error) {
	User, err := u.UserRepository.FindUserByName(userName)
	if err!=nil{
		//have not the user
		return false,err
	}
	return ValidatePassword(pwd,User.HashPassword)
}