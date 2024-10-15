package env

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	MysqlDB *gorm.DB
	Gin     *gin.Engine
	RedisDB *redis.Client
)
