package main

import (
	"log"
	"os"
	"web/env"
)

// 配置文件默认路径
const defaultConfigFile = "./config/application.yaml"

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
