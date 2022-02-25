package entity

import "go_dev/day17/demo/src/share/pb"

type User struct {
	Id 		int32 		`json:"id" db:"id"`
	Name 	string 		`json:"name" db:"name"`
	Address string		`json:"address" db:"address"`
	Phone   string		`json:"phone" db:"phone"`
}

func (u *User) ToProtoUser() *pb.User {
	return &pb.User{
		Id: 	u.Id ,
		Name: 	u.Name ,
		Address:u.Address,
		Phone:  u.Phone  ,
	}
}