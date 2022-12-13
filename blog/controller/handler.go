package controller

import (
	"blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CategoryHandler(ctx *gin.Context) {
	categoryList, err := service.GetAllCateGategoryList()
	if err!=nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"massage":"server err ,pleace try agin later",
		})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{
		"categoryList":categoryList,
		"massage":"ok",
	})
	return
}