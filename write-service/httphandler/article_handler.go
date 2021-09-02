package httphandler

import (
	"article-web-service/shared/domains"
	dto "article-web-service/write-service/dto/request"
	"article-web-service/write-service/service"
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

func (h *articleHandler) Store(c *gin.Context) {
	var reqBody dto.NewArticleRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		rs, err := domains.FormatValidationErrors(h.AppContext.Localizer, err)
		if err != nil {
			domains.InternalServerError(c, err.Error())
			return
		}

		domains.BadRequest(c, rs)
		return
	}

	srv := service.NewStoreNewArticleService(h.AppContext, reqBody)
	if err := srv.Run(); err != nil {
		domains.InternalServerError(c, err.Error())
		return
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}
