package route

import (
	"golearn/internal/app/service"
	"golearn/internal/pkg/model"
	"golearn/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

// 创建用户
func CreateUser(c *gin.Context) {
	var param model.User
	c.BindJSON(&param)
	service.CreateUser(&param)
	response.Success()
}

// 更新用户
func UpdateUser(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		response.Failed(&response.INVALID_PARAM)
		return
	}
	var param model.User
	c.BindJSON(&param)
	service.UpdateUser(id, &param)
	response.Success()
}

// 删除用户
func DeleteUser(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		response.Failed(&response.INVALID_PARAM)
		return
	}
	service.DeleteUser(id)
	response.Success()
}

// 查询用户
func SelectUser(c *gin.Context) {
	var param model.User
	c.BindJSON(&param)
	response.SuccessReturn(service.FindUserByNameLike(param.Name))
}
