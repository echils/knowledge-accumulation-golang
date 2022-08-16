package main

import (
	"fmt"
)

//定义函数类型
type doubleIntFunc func(i1 int, i2 int) int

//定义方法使用函数类型
func testFunc(dif doubleIntFunc, max int, min int) int {
	return dif(max, min)
}

func min(i1 int, i2 int) int {
	if i1 < i2 {
		return i1
	}
	return i2
}

//定义返回值时直接定义变量，这样return就可以什么也不屑
func cal(i1 int, i2 int) (sum int, sub int) {
	sum = i1 + i2
	sub = i1 - i2
	return
}

func main() {

	var testMin = min

	fmt.Println(testMin(50, 40))
	fmt.Println(testFunc(testMin, 50, 40))

	//忽略sub值时可以用_代替
	sum, _ := cal(50, 40)
	fmt.Println(sum)
}
