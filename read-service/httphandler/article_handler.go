package httphandler

import (
	"article-web-service/read-service/service"
	"article-web-service/shared/domains"
	"net/http"

	"github.com/gin-gonic/gin"
)

type articleHandler struct {
	AppContext *domains.Context
}

func NewArticleHandler(appCtx *domains.Context) *articleHandler {
	return &articleHandler{
		AppContext: appCtx,
	}
}

func (h *articleHandler) Index(c *gin.Context) {
	query := c.Query("query")
	author := c.Query("author")

	srv := service.NewListArticleService(h.AppContext, query, author)
	articles, err := srv.Run()
	if err != nil {
		domains.InternalServerError(c, err.Error())
		return
	}

	domains.SuccessResponse(c, http.StatusOK, articles, nil)
}
