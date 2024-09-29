package main

import (
	"fmt"
	"golearn/pkg/utils"
)

func main() {

	fmt.Println(utils.SnowflakeID())

	fmt.Println(utils.Random32UUID())
}
