package controller

import (
	"web/model"
	"web/service"

	"github.com/gin-gonic/gin"
)

// 配置用户管理路由
func RegisterUserControllerRoute(e *gin.Engine) {
	group := e.Group("/user")
	group.POST("/create", CreateUser)
	group.PUT("/:id/update", UpdateUser)
	group.DELETE("/:id/delete", DeleteUser)
	group.GET("/list", SelectUser)
}

// 创建用户
func CreateUser(c *gin.Context) {
	var param model.User
	c.BindJSON(&param)
	service.CreateUser(&param)
	model.InsertSuccess(c, nil)
}

// 更新用户
func UpdateUser(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		model.Failed(c, "无效参数ID")
		return
	}
	var param model.User
	c.BindJSON(&param)
	service.UpdateUser(id, &param)
	model.UpdateSuccess(c, nil)
}

// 删除用户
func DeleteUser(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		model.Failed(c, "无效参数ID")
		return
	}
	service.DeleteUser(id)
	model.DeleteSuccess(c, nil)
}

// 查询用户
func SelectUser(c *gin.Context) {
	var param model.User
	c.BindJSON(&param)
	model.Success(c, service.FindUserByNameLike(param.Name))
}
