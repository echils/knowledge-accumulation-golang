package main

import (
	"web/pkg/setting"
)

// 环境变量相对路径
const configFile = "./config/app.yaml"

func main() {

	//项目环境配置引导
	b, _ := setting.Bootstrap(configFile)
	if !b {
		panic("fatal error when bootstrap")
	}

}
