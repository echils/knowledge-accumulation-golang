package controller

import (
	"web/model"

	"github.com/gin-gonic/gin"
)

// 配置用户管理路由
func RegisterUserControllerRoute(e *gin.Engine) {
	group := e.Group("/user")
	group.POST("/create", CreateUser)
	group.PUT("/update", UpdateUser)
	group.DELETE("/delete", DeleteUser)
	group.GET("/select", SelectUser)
}

// 创建用户
func CreateUser(c *gin.Context) {

}

// 更新用户
func UpdateUser(c *gin.Context) {

}

// 删除用户
func DeleteUser(c *gin.Context) {

}

// 查询用户
func SelectUser(c *gin.Context) {
	model.Success(c, model.User{ID: 1, Name: "Java", Age: 8})
}
