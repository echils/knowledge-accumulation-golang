package app

import (
	"context"
	"errors"
	"fmt"
	"golearn/internal/app/controller"
	"golearn/internal/app/env"
	"golearn/internal/pkg/response"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	logC "github.com/lestrrat-go/file-rotatelogs"
	"github.com/redis/go-redis/v9"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 启动服务
func Run(path string) (config Config, err error) {

	//读取配置文件
	config, err = loadConfigFile(path)
	if err != nil {
		panic(fmt.Errorf("fatal error when load config file: %w", err))
	}

	//连接MySQL
	env.MysqlDB, err = connectMysql(config.MySQL)
	if err != nil {
		panic(fmt.Errorf("fatal error when connect mysql: %w", err))
	}

	//连接Redis
	env.RedisDB, err = connectRedis(config.Redis)
	if err != nil {
		panic(fmt.Errorf("fatal error when connect redis: %w", err))
	}

	//启动服务
	gin.ForceConsoleColor()
	gin.SetMode(config.Server.Mode)

	engine := gin.Default()
	engine.Use(loggerToFile(config), globalException())

	//加载路由
	loadRouteConfig(engine)

	port := "8000"
	if config.Server.Port != 0 {
		port = strconv.Itoa(config.Server.Port)
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: engine,
	}

	log.Printf("Start application with profile [%s] and port [%s] \n", config.Server.Mode, port)
	err = server.ListenAndServe()
	if err != nil {
		panic(fmt.Errorf("start application occur error when listen sever: %v", err))
	}
	return config, nil
}

// 加载路由
func loadRouteConfig(e *gin.Engine) {
	group := e.Group("/api/v1")
	group.POST("/user/create", controller.CreateUser)
	group.PUT("/user/:id/update", controller.UpdateUser)
	group.DELETE("/user/:id/delete", controller.DeleteUser)
	group.GET("/user/list", controller.SelectUser)
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
func connectMysql(config MySQL) (db *gorm.DB, e error) {
	dsn := fmt.Sprintf("%s:%s@%s", config.Username, config.Password, config.Url)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, errors.New("Failed to connect database: " + err.Error())
	}
	return db, nil
}

// 连接Redis
func connectRedis(config Redis) (redisDB *redis.Client, e error) {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Username: config.Username,
		Password: config.Password,
		DB:       config.Database,
	})
	_, err := redisDB.Ping(context.Background()).Result()
	if err != nil {
		return nil, errors.New("Redis connect error" + err.Error())
	}
	return redisDB, nil
}

// 全局异常处理
func globalException() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Fatalln("System occur error: ", fmt.Sprintf("%v", err))
				response.Failed(&response.DEFAULT_ERROR)
				return
			}
		}()
		c.Next()
	}
}

// 日志配置
func loggerToFile(config Config) gin.HandlerFunc {

	logger := logrus.New()
	logName := config.Server.Name + ".log"
	fileName := path.Join(config.Log.Dir, logName)
	src, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		panic("Log open file error: " + err.Error())
	}
	logger.Out = src
	logger.SetLevel(logrus.Level(config.Log.Level))

	logWriter, err := logC.New(
		fileName+".%Y%m%d.log",
		logC.WithLinkName(fileName),
		logC.WithMaxAge(7*24*time.Hour),
		logC.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		panic("Log rotate config error: " + err.Error())
	}

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	hook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.AddHook(hook)

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		consume := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": consume,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
		}).Info()
	}

}

// 服务环境信息
type Config struct {
	//服务信息
	Server Server `yaml:"server"`
	//日志信息
	Log Log `yaml:"log"`
	//MySQL信息
	MySQL MySQL `yaml:"mysql"`
	//Redis信息
	Redis Redis `yaml:"redis"`
}

// 获取Log配置
type Log struct {
	Level int    `yaml:"level"`
	Dir   string `yaml:"dir"`
}

// 获取Server配置
type Server struct {
	Name string `yaml:"name"`
	Port int    `yaml:"port"`
	Mode string `yaml:"mode"`
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
