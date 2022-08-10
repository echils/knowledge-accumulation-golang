package main

import "fmt"

func init() {
	fmt.Println("========================")
	fmt.Println("Welcome to Golang world")
}

func main() {

	var a = 1
	b := 5

	fmt.Print(add(a, b))

}

func add(param ...int) int {
	if len(param) == 0 {
		return 0
	}
	result := 0
	for i := range param {
		result += param[i]
	}
	return result
}
