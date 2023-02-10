package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm_connect/model"
	"gorm_connect/service"
)

// 添加用户
func AddUser(ctx *gin.Context) {
	var u model.User
	if err := ctx.ShouldBind(&u); err != nil {
		panic(err)
	}
	if err := service.AddUser(u.Username, u.Password); err != nil {
		fmt.Printf("errorAddUser:%#v", err)
		ctx.JSON(http.StatusOK, gin.H{
			"massage": err.Error(),
		})
		return
	}
	fmt.Printf("user:%#v", &u)
	ctx.JSON(http.StatusOK, gin.H{
		"massage": "status ok",
	})
}

// login
func Login(ctx *gin.Context) {
	var u model.User
	if err := ctx.ShouldBind(&u); err != nil {
		panic(err)
	}
	if err := service.Login(u.Username, u.Password); err != nil {
		fmt.Printf("errorlogin:%#v", err)
		ctx.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	fmt.Printf("user:%#v", &u)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "登陆成功",
	})
}

//change password

