package handler

import (
	"context"
	"user/domain/model"
	"user/domain/service"
	user "user/proto/user"
)

type User struct{
	UserdDataService service.IUserDataService
}
	/**Registry user
	   1.generate the user
	   2.add the new user to Repository.Add
	   3.detection the add
	 */
	func(u *User)Register(ctx context.Context, req *user.UserRegisterRequest,rep *user.UserRegisterResponse) error{
			userRegister :=&model.User{
				UserName:req.UserName,
				FirstName:req.FirstName,
				HashPassword:req.Pwd,
			}
		userID, err := u.UserdDataService.AddUser(userRegister)
		if err!=nil{
			return err
		}
		rep.Message="新用户添加成功"+string(userID)
		return nil
	}
	//login
	func(u *User)Login(ctx context.Context, req *user.UserLoginRequest, rep *user.UserLoginResponse) error{
		isOk, err := u.UserdDataService.CheckPwd(req.UserName, req.Pwd)
		if err!=nil{
			return err
		}
		rep.IsSuccess=isOk
		return nil
	}
	/**query user info
	   1.find user info
	   2.generate User for response
	 */
	func(u *User)GetUseInfo(ctx context.Context,req *user.UserInfoRequest, rep *user.UserInfoResponse) error{
		userGetinfo, err := u.UserdDataService.FindUserByName(req.UserName)
		if err!=nil{
			return err
		}
		rep = UserforResponse(userGetinfo)
		return nil
 	}

	//user for response
	func UserforResponse(u *model.User) *user.UserInfoResponse  {
		 res :=&user.UserInfoResponse{}
		 res.UserName=u.UserName
		 res.FirstName=u.FirstName
		 res.UserId=u.ID
		 return res
	}

