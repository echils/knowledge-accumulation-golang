package env

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/redis/go-redis/v9"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Config  Env
	MysqlDB *gorm.DB
	RedisDB *redis.Client
	Engine  *gin.Engine
)

// 服务环境信息
type Env struct {
	//日志信息
	Log Log `yaml:"log"`
	//服务信息
	Server Server `yaml:"server"`
	//MySQL信息
	MySQL MySQL `yaml:"mysql"`
	//Redis信息
	Redis Redis `yaml:"redis"`
	//MinIO信息
	MinIO MinIO `yaml:"minio"`
}

// 获取Log配置
type Log struct {
	Level   string `yaml:"level"`
	LogFile string `yaml:"file"`
}

// 获取Server配置
type Server struct {
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

// 获取MinIO配置
type MinIO struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// 启动Web服务
func Run() {
	gin.ForceConsoleColor()
	gin.SetMode(Config.Server.Profile)
	Engine = gin.Default()
	if Config.Server.Port != 0 {
		Engine.Run(":" + strconv.Itoa(Config.Server.Port))
	} else {
		Engine.Run()
	}

	log.Println(fmt.Printf("Started application with profile %s and port %d", Config.Server.Profile, Config.Server.Port))
}

// 配置环境引导
func Bootstrap(path string) (s bool, e Env) {

	//加载配置文件
	e, r := loadConfigFile(path)
	if r != nil {
		log.Println(fmt.Errorf("fatal error when load config file: %w", r))
		return false, e
	}

	//连接MySQL
	connectMysql(e.MySQL)
	//连接Redis
	connectRedis(e.Redis)
	//连接MinIO
	// connectMinIO(e.MinIO)

	Config = e
	return true, e
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

// 连接MySQL
func connectMysql(config MySQL) {
	dsn := fmt.Sprintf("%s:%s@%s", config.Username, config.Password, config.Url)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database: " + err.Error())
	}
	MysqlDB = db
}

// 连接Redis
func connectRedis(config Redis) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Username: config.Username,
		Password: config.Password,
		DB:       config.Database,
	})
	RedisDB = rdb
}

// 连接MinIO
func connectMinIO(config MinIO) {
	minioClient, err := minio.New(fmt.Sprintf("%s:%d", config.Host, config.Port), &minio.Options{
		Creds: credentials.NewStaticV4(config.Username, config.Password, ""),
	})
	if err != nil {
		panic("Failed to connect minio: " + err.Error())
	}
	fmt.Println(minioClient)
}
