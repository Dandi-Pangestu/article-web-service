package domains

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
)

type Context struct {
	R           *gin.Engine
	DB          *gorm.DB
	Localizer   *i18n.Localizer
	RedisClient *redis.Client
	ESClient    *elastic.Client
}
