package main

import (
	"golearn/internal/app"
	"log"
	"os"
)

// 配置文件默认路径
const defaultConfigFile = "./configs/app.yaml"

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
	app.Run(configFile)
}
