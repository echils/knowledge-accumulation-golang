package env

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"time"
	"web/controller"
	"web/model"

	"github.com/gin-gonic/gin"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/redis/go-redis/v9"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Env     Config
	MysqlDB *gorm.DB
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
func Run(configFile string) {

	//项目环境配置引导
	b := bootstrap(configFile)
	if !b {
		panic("Panic error when bootstrap")
	}

	//启动Gin
	gin.ForceConsoleColor()
	gin.SetMode(Env.Server.Profile)
	e := gin.Default()
	e.Use(loggerToFile(), globalException())

	//路由映射
	loadRoutes(e)

	if Env.Server.Port != 0 {
		log.Printf("Started application with profile [%s] and port [%d] \n", Env.Server.Profile, Env.Server.Port)
		e.Run(":" + strconv.Itoa(Env.Server.Port))
	} else {
		log.Printf("Started application with profile [%s] and port [8080]", Env.Server.Profile)
		e.Run()
	}

}

// 加载路由
func loadRoutes(e *gin.Engine) {
	controller.RegisterUserControllerRoute(e)
}

// 全局异常处理
func globalException() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				model.Failed(c, "系统响应异常")
				return
			}
		}()
		c.Next()
	}
}

// 配置环境引导
func bootstrap(path string) (s bool) {

	//加载配置文件
	err := loadConfigFile(path)
	if err != nil {
		log.Println(fmt.Errorf("fatal error when load config file: %w", err))
		return false
	}

	//连接MySQL
	connectMysql(Env.MySQL)
	//连接Redis
	connectRedis(Env.Redis)

	return true
}

// 加载配置文件
func loadConfigFile(path string) (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = errors.New(fmt.Sprint(err))
		}
	}()
	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln("OS read config file failed: ", path)
		return err
	}
	err2 := yaml.Unmarshal(b, &Env)
	if err2 != nil {
		log.Fatalln("Yaml unmarshal config file failed: ", path)
		return err2
	}
	return nil
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
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Username: config.Username,
		Password: config.Password,
		DB:       config.Database,
	})
}

// 日志配置
func loggerToFile() gin.HandlerFunc {

	logger := logrus.New()
	logName := Env.Server.Name + ".log"
	fileName := path.Join(Env.Log.LogFile, logName)
	src, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		panic("Log open file error: " + err.Error())
	}
	logger.Out = src
	logger.SetLevel(logrus.Level(Env.Log.Level))

	logWriter, err := rotateLogs.New(
		fileName+".%Y%m%d.log",
		rotateLogs.WithLinkName(fileName),
		rotateLogs.WithMaxAge(7*24*time.Hour),
		rotateLogs.WithRotationTime(24*time.Hour),
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
