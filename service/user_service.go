package service

import (
	"context"
	"web/dao"
	"web/env"
	"web/model"
)

// 创建用户
func CreateUser(user *model.User) {
	user.Id = model.RandomUUID()
	if len(user.Name) == 0 {
		panic("用户姓名不能为空")
	}
	if user.Age < 18 {
		panic("用户未满18岁")
	}
	dao.InsertUser(user)
	env.RedisDB.HMSet(context.Background(), "user", user.Id, user.Name)
}

// 更新用户
func UpdateUser(id string, user *model.User) {

	dataUser := dao.FindUserById(id)
	if len(user.Name) == 0 {
		panic("用户姓名不能为空")
	}
	dataUser.Name = user.Name
	if user.Age < 18 {
		panic("用户未满18岁")
	}
	dataUser.Age = user.Age
	dao.UpdateUser(&dataUser)
	env.RedisDB.HMSet(context.Background(), "user", user.Id, user.Name)
}

// 删除用户
func DeleteUser(id string) {
	dao.DeleteUser(id)
	env.RedisDB.HDel(context.Background(), "user", id)

}

// 通过用户姓名模糊查询
func FindUserByNameLike(name string) (users []model.User) {
	if len(name) == 0 {
		panic("参数异常")
	}
	return dao.FindUserByNameLike(name)
}
