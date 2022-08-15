package main

import "fmt"

func main() {

	//定义通道
	c := make(chan int)

	//异步执行函数
	go test(c)

	//主线程阻塞获取结果
	result := <-c

	fmt.Println(result)

	close(c)

}

//定义函数将结果通过通道复制给result
func test(result chan int) {
	var value int
	for i := 0; i <= 1000; i++ {
		value += i
	}
	result <- value
}
