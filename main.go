package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	// 通过net/http实现后端web服务
	http.HandleFunc("/", sayHello)

	//监听8080端口
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("http server failed,error:%v\n", err)
		return
	}

}

func sayHello(writer http.ResponseWriter, request *http.Request) {
	b, err := os.ReadFile("./data.txt")
	if err != nil {
		fmt.Fprintln(writer, "系统响应异常")
	}
	fmt.Fprintln(writer, string(b))
}
