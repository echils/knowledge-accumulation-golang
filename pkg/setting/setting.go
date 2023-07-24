package setting

import (
	"errors"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// 服务环境信息
type Env struct {
	//服务信息
	Server Server `yaml:"server"`
	//MySQL信息
	MySQL MySQL `yaml:"mysql"`
	//Redis信息
	Redis Redis `yaml:"redis"`
	//MinIO信息
	MinIO MinIO `yaml:"minio"`
}

// 获取Server配置
type Server struct {
	Port int `yaml:"port"`
}

// 获取MySQL配置
type MySQL struct {
	Url      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// 获取Redis配置
type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database int    `yaml:"database"`
}

// 获取MinIO配置
type MinIO struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// 配置环境引导
func Bootstrap(path string) (s bool, e Env) {

	//加载配置文件
	e, r := loadConfigFile(path)
	if r != nil {
		log.Fatalln(fmt.Errorf("fatal error when load config file: %w", r))
		return false, e
	}

	//连接MySQL
	connectMysql(e.MySQL)

	//连接Redis
	connectRedis(e.Redis)

	//连接MinIO
	connectMinIO(e.MinIO)

	return true, e
}

// 连接MySQL
func connectMysql(mysql MySQL) {
	
}

// 连接Redis
func connectRedis(redis Redis) {

}

// 连接MinIO
func connectMinIO(minio MinIO) {

}

// 加载配置文件
func loadConfigFile(path string) (env Env, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = errors.New(fmt.Sprint(err))
		}
	}()
	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln("OS read config file failed: ", path)
		return env, err
	}
	err2 := yaml.Unmarshal(b, &env)
	if err2 != nil {
		log.Fatalln("Yaml unmarshal config file failed: ", path)
		return env, err2
	}
	return env, nil
}
