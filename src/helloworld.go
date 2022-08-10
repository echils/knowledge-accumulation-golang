package main

import "fmt"

//定义枚举
const (
	GO     = "Golang"
	JAVA   = "Java"
	PYTHON = "Python"
	OTHER  = "Other"
)

//初始方法，先于main函数执行，可以有多个init函数
func init() {
	fmt.Println("========================")
	fmt.Println("Welcome to the programmer's world!")
}
func init() {
	fmt.Println("I have mastered the "+JAVA+" language", "and start to learn "+GO+" language at now")
	fmt.Println("========================")
}

func main() {

	//定义变量和常量，iota等于0可以用它来定义枚举，会自动+1
	var a = 1.9
	b := 5
	const c = iota
	var d, e int
	d, e = 4, 10
	fmt.Println(sum(int(a), b, c, d, e))
	fmt.Println("========================")

	//定义数组
	var project [3]string
	project[0] = GO
	project[1] = PYTHON
	project[2] = JAVA
	for i := 0; i < len(project); i++ {
		fmt.Print(project[i], " ")
	}
	var language = [...]string{JAVA, GO, PYTHON}
	for i := range language {
		fmt.Println(language[i])
	}
	fmt.Println("========================")

	//定义切片，类似Java的ArrayList
	//定义空切片
	var content []string
	fmt.Println(content, len(content), cap(content), content == nil)
	//定义一个初始长度为3的切片，容量缺省，默认等于初始长度
	content = make([]string, 3)
	fmt.Println(content, len(content), cap(content), content == nil)
	content[0] = "0"
	content[1] = "1"
	content[2] = "2"
	//content[3] = "3" 大于容量是包越界错误
	fmt.Println(content, len(content), cap(content), content == nil)
	//使用append函数会自动2倍扩容
	content = append(content, "1")
	fmt.Println(content, len(content), cap(content), content == nil)
	for i, s := range content {
		fmt.Printf("index:%v,value:%s\n", i, s)
	}
	fmt.Println("========================")

	//定义Map
	languageMap := make(map[int]string)
	languageMap[0] = JAVA
	languageMap[1] = GO
	languageMap[2] = PYTHON
	languageMap[3] = OTHER
	//map遍历时无序
	for i := range languageMap {
		fmt.Print(languageMap[i], " ")
	}
	fmt.Println()
	delete(languageMap, 3)
	for i, s := range languageMap {
		fmt.Printf("index:%v,value:%s\n", i, s)
	}
	fmt.Println()
	fmt.Println("========================")
}

func sum(param ...int) int {
	if len(param) == 0 {
		return 0
	}
	result := 0
	for i := range param {
		result += param[i]
	}
	return result
}
