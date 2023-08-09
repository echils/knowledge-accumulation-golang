package main

import (
	"log"
	"os"
	"web/env"
)

// 环境变量相对路径
const defaultConfigFile = "./conf/config.yaml"

func main() {

	//加载配置文件
	configFile := defaultConfigFile
	if len(os.Args) > 2 {
		log.Println("Load specified config file", os.Args[1])
		configFile = os.Args[1]
	} else {
		log.Println("Load default config file", defaultConfigFile)
	}

	//启动服务
	env.Run(configFile)
}
