package controller

import (
	"golearn/internal/app/service"
	"golearn/internal/pkg/model"
	"golearn/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

// 创建用户
func CreateUser(gin *gin.Context) {
	var param model.User
	gin.BindJSON(&param)
	service.CreateUser(&param)
	gin.JSON(200, response.Success())
}

// 更新用户
func UpdateUser(gin *gin.Context) {
	id, ok := gin.Params.Get("id")
	if !ok {
		response.Failed(&response.INVALID_PARAM)
		return
	}
	var param model.User
	gin.BindJSON(&param)
	service.UpdateUser(id, &param)
	gin.JSON(200, response.Success())
}

// 删除用户
func DeleteUser(gin *gin.Context) {
	id, ok := gin.Params.Get("id")
	if !ok {
		response.Failed(&response.INVALID_PARAM)
		return
	}
	service.DeleteUser(id)
	gin.JSON(200, response.Success())
}

// 查询用户
func SelectUser(gin *gin.Context) {
	var param model.User
	gin.BindJSON(&param)
	gin.JSON(200, response.SuccessReturn(service.FindUserByNameLike(param.Name)))
}
