package controller

import (
	"golearn/internal/app/service"
	"golearn/internal/pkg/model"
	"golearn/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

// @Summary 创建用户
// @tags 用户管理
// @Accept application/json
// @Produce application/json
// @Param name body string true "用户名称"
// @Success 200 {object} response.ResultData "成功"
// @Router /api/v1/user/create [post]
func CreateUser(gin *gin.Context) {
	var param model.User
	gin.BindJSON(&param)
	service.CreateUser(&param)
	response.Success(gin)
}

// 更新用户
func UpdateUser(gin *gin.Context) {
	id, ok := gin.Params.Get("id")
	if !ok {
		response.Failed(gin, &response.INVALID_PARAM)
		return
	}
	var param model.User
	gin.BindJSON(&param)
	service.UpdateUser(id, &param)
	response.Success(gin)
}

// 删除用户
func DeleteUser(gin *gin.Context) {
	id, ok := gin.Params.Get("id")
	if !ok {
		response.Failed(gin, &response.INVALID_PARAM)
		return
	}
	service.DeleteUser(id)
	response.Success(gin)
}

// 查询用户
func SelectUser(gin *gin.Context) {
	var param model.User
	gin.BindJSON(&param)
	response.SuccessReturn(gin, service.FindUserByNameLike(param.Name))
}
