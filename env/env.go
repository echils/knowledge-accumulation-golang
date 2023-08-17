package env

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	MysqlDB *gorm.DB
	Gin     *gin.Engine
	RedisDB *redis.Client
)

// 服务环境信息
type Config struct {
	//日志信息
	Log Log `yaml:"log"`
	//服务信息
	Server Server `yaml:"server"`
	//MySQL信息
	MySQL MySQL `yaml:"mysql"`
	//Redis信息
	Redis Redis `yaml:"redis"`
}

// 获取Log配置
type Log struct {
	Level   int    `yaml:"level"`
	LogFile string `yaml:"file"`
}

// 获取Server配置
type Server struct {
	Name    string `yaml:"name"`
	Port    int    `yaml:"port"`
	Profile string `yaml:"profile"`
	LogFile string `yaml:"log-file"`
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

// 启动Web服务
func Bootstrap(configFile string) (config Config) {

	//项目环境配置引导
	env, err := loadConfigFile(configFile)
	if err != nil {
		panic(fmt.Errorf("fatal error when load config file: %w", err))
	}

	//连接MySQL
	err = connectMysql(env.MySQL)
	if err != nil {
		panic(fmt.Errorf("fatal error when connect mysql: %w", err))
	}

	//连接Redis
	err = connectRedis(env.Redis)
	if err != nil {
		panic(fmt.Errorf("fatal error when connect redis: %w", err))
	}
	return env
}

// 加载配置文件
func loadConfigFile(path string) (config Config, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = errors.New(fmt.Sprint(err))
		}
	}()
	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln("OS read config file failed: ", path)
		return config, err
	}
	err2 := yaml.Unmarshal(b, &config)
	if err2 != nil {
		log.Fatalln("Yaml unmarshal config file failed: ", path)
		return config, err2
	}
	return config, nil
}

// 连接MySQL
func connectMysql(config MySQL) (e error) {
	dsn := fmt.Sprintf("%s:%s@%s", config.Username, config.Password, config.Url)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return errors.New("Failed to connect database: " + err.Error())
	}
	MysqlDB = db
	return nil
}

// 连接Redis
func connectRedis(config Redis) (e error) {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Username: config.Username,
		Password: config.Password,
		DB:       config.Database,
	})
	_, err := RedisDB.Ping(context.Background()).Result()
	if err != nil {
		return errors.New("Redis connect error" + err.Error())
	}
	return nil
}
