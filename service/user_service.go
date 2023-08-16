package service

import (
	"errors"
	"web/dao"
	"web/model"
)

// 创建用户
func CreateUser(user *model.User) (e error) {
	user.Id = model.RandomUUID()
	if len(user.Name) == 0 {
		return errors.New("用户姓名不能为空")
	}
	if user.Age < 18 {
		return errors.New("用户未满18岁")
	}
	dao.InsertUser(user)
	return nil
}

// 更新用户
func UpdateUser(id string, user *model.User) (e error) {

	dataUser := dao.FindUserById(id)
	if len(user.Name) == 0 {
		return errors.New("用户姓名不能为空")
	}
	dataUser.Name = user.Name
	if user.Age < 18 {
		return errors.New("用户未满18岁")
	}
	dataUser.Age = user.Age
	dao.UpdateUser(&dataUser)
	return nil
}

// 删除用户
func DeleteUser(id string) (e error) {
	dao.DeleteUser(id)
	return nil
}

// 通过用户姓名模糊查询
func FindUserByNameLike(name string, user *model.User) (users []model.User, e error) {
	if len(name) == 0 {
		return nil, errors.New("参数异常")
	}
	return dao.FindUserByNameLike(name), nil
}
