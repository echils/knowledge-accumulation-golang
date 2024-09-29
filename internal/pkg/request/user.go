package request

type UserListDTO struct {
	Ids           []string `json:"ids"`
	Name          string   `json:"name"`
	DepartmentIds []string `json:"departmentIds"`
}

type UserAddDTO struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserUpdateDTO struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserDeleteDTO struct {
	Ids []string `json:"ids"`
}
