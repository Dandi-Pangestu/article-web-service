package router

import (
	"article-web-service/read-service/httphandler"
	"article-web-service/shared/domains"
)

func Configure(c *domains.Context) {
	v1 := c.R.Group("/v1")
	{
		articleRouter := v1.Group("/articles")
		{
			handler := httphandler.NewArticleHandler(c)
			articleRouter.GET("", handler.Index)
		}
	}
}
