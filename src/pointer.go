package main

import "fmt"

func main() {

	variablePointerTest()
	fmt.Println("====================")
	doublePointerTest()
	fmt.Println("====================")
	content := "Golang"
	fmt.Println(methodPointerTest(&content))

}

//变量指针
func variablePointerTest() {

	//定义空指针
	var ptr *int
	fmt.Println(ptr)
	a := 5

	//指针赋值
	ptr = &a
	fmt.Println(ptr)

	//获取指针对应的值
	fmt.Println(*ptr)

}

//双重指针
func doublePointerTest() {

	var content = "Pointer"
	var ptr *string
	var ptrDouble **string

	fmt.Println(ptr)
	ptr = &content
	fmt.Println(ptr)
	fmt.Println(*ptr)
	ptrDouble = &ptr
	fmt.Println(ptrDouble)
	fmt.Println(**ptrDouble)

}

//函数指针
func methodPointerTest(ptr *string) string {
	return *ptr
}
