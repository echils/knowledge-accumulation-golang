package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(User{1, "张三", 18, false, "GO"})
	fmt.Println("=====================")

	user := User{
		id:         2,
		name:       "李四",
		age:        18,
		department: "Java",
	}

	fmt.Println(user.toString())
	fmt.Println(user.name)
	fmt.Println(user.department)
	user.age = 19
	fmt.Println(user.toString())
	fmt.Println("=====================")

	//值传递，更改不影响原数据，GO默认值传递
	valuePassed(user)
	fmt.Println(user.toString())

	fmt.Println("=====================")
	//引用传递，更改影响原数据
	referencePassed(&user)
	fmt.Println(user.toString())

}

//定义用户结构体
type User struct {
	id         uint
	name       string
	age        int
	male       bool
	department string
}

//定义User结构体特有的方法
func (user User) toString() string {
	return "{id=" + fmt.Sprintf("%v", user.id) + ", name=" + user.name + ", age=" +
		fmt.Sprintf("%v", user.age) + ", male=" + strconv.FormatBool(user.male) + ", department=" + user.department + "}"
}

//方法值传递
func valuePassed(user User) {
	user.age = 100
}

//方法引用传递
func referencePassed(user *User) {
	user.age = 100
}
