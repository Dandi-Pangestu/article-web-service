package middleware

import (
	"article-web-service/shared/domains"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func Localizer(appContext *domains.Context, bundle *i18n.Bundle) gin.HandlerFunc {
	return func(c *gin.Context) {
		accept := c.GetHeader("Accept-Language")
		appContext.Localizer = i18n.NewLocalizer(bundle, accept)
		c.Set("localizer", appContext.Localizer)
		c.Set("lang", accept)
	}
}
