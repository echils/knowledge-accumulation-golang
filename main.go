package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"time"
	"web/controller"
	"web/env"
	"web/model"

	"github.com/gin-gonic/gin"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
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

	//引导配置
	env := env.Bootstrap(configFile)

	//启动服务
	start(env)

}

func start(config env.Config) {

	gin.ForceConsoleColor()
	gin.SetMode(config.Server.Profile)
	engine := gin.Default()
	engine.Use(loggerToFile(config), globalException())

	//路由映射
	loadRoutes(engine)

	//启动服务
	if config.Server.Port != 0 {
		log.Printf("Started application with profile [%s] and port [%d] \n", config.Server.Profile, config.Server.Port)
		engine.Run(":" + strconv.Itoa(config.Server.Port))
	} else {
		log.Printf("Started application with profile [%s] and port [8080]", config.Server.Profile)
		engine.Run()
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
				model.Failed(c, fmt.Sprintf("%v", err))
				return
			}
		}()
		c.Next()
	}
}

// 日志配置
func loggerToFile(config env.Config) gin.HandlerFunc {

	logger := logrus.New()
	logName := config.Server.Name + ".log"
	fileName := path.Join(config.Log.LogFile, logName)
	src, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		panic("Log open file error: " + err.Error())
	}
	logger.Out = src
	logger.SetLevel(logrus.Level(config.Log.Level))

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
