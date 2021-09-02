package router

import (
	"article-web-service/shared/domains"
	"article-web-service/write-service/httphandler"
)

func Configure(c *domains.Context) {
	v1 := c.R.Group("/v1")
	{
		articleRouter := v1.Group("/articles")
		{
			handler := httphandler.NewArticleHandler(c)
			articleRouter.POST("", handler.Store)
		}
	}
}
