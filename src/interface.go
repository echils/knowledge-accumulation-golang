package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	iphone := Phone{"Apple"}
	fmt.Println(Message(iphone).send())
	fmt.Println(Message(iphone).receive())
	email := Email{"网易"}

	send, err := Message(email).send()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(send)
	}
	fmt.Println(Message(email).receive())
}

//定义消息接口
type Message interface {
	//发送
	send() (string, error)
	//接收
	receive() (string, error)
}

type Phone struct {
	//品牌
	vendor string
}

func (p Phone) send() (string, error) {
	panic("手机没电了，发送失败")
}

func (p Phone) receive() (string, error) {
	return p.vendor + "手机收到一条消息", nil
}

type Email struct {
	//账号
	account string
}

func (e Email) send() (string, error) {
	return e.account + "邮箱发送一条消息", errors.New(e.account + "邮箱发送失败")
}

func (e Email) receive() (string, error) {
	return e.account + "邮箱收到一条消息", nil
}
