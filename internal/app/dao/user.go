package dao

import (
	"golearn/internal/app"
	"golearn/internal/pkg/model"
)

// 插入用户至数据库
func InsertUser(user *model.User) {
	app.MysqlDB.Create(user)
}

// 更新用户至数据库
func UpdateUser(user *model.User) {
	app.MysqlDB.Updates(user)
}

// 通过用户ID查询
func FindUserById(id string) (user model.User) {
	user = model.User{Id: id}
	app.MysqlDB.First(&user)
	return user
}

// 通过用户ID删除
func DeleteUser(id string) (user model.User) {
	user = model.User{Id: id}
	app.MysqlDB.Delete(&user)
	return user
}

// 通过用户姓名模糊查询
func FindUserByNameLike(name string) (users []model.User) {
	app.MysqlDB.Where("name like ?", "%"+name+"%").Find(&users)
	return users
}
