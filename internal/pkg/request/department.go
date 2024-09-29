package request

type DepartmentListDTO struct {
	Ids  []string `json:"ids"`
	Name string   `json:"name"`
}

type DepartmentAddDTO struct {
	Name string `json:"name"`
}

type DepartmentUpdateDTO struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type DepartmentDeleteDTO struct {
	Ids []string `json:"ids"`
}
