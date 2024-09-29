package model

// 用户信息
type User struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Age          int    `json:"age"`
	DepartmentId string `json:"departmentId"`
}

// 用户视图
type UserVo struct {
	User
	DepartmentInfo Department `json:"departmentInfo"`
}
