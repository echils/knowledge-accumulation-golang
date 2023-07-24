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

	//项目环境配置引导
	b, e := env.Bootstrap(configFile)
	if !b {
		panic("Panic error when bootstrap")
	}

	//日志写入文件
	logFile, err := os.OpenFile(e.Log.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic("Panic error when open log file: " + err.Error())
	}
	log.SetOutput(logFile)

	//启动服务
	env.Run()
}
