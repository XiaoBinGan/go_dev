package handler

import "go_dev/day17/demo/src/user-srv/entity"

func init() {
	user := entity.User.ToProtoUser()
}