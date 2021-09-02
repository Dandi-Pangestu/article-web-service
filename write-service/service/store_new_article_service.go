package service

import (
	"article-web-service/shared/domains"
	log "article-web-service/shared/log/app"
	"article-web-service/shared/utils"
	dto "article-web-service/write-service/dto/request"
	"article-web-service/write-service/models"

	"github.com/hashicorp/go.net/context"
	"github.com/sirupsen/logrus"
)

type storeNewArticleService struct {
	AppContext *domains.Context
	ReqBody    dto.NewArticleRequest
}

func NewStoreNewArticleService(appCtx *domains.Context, reqBody dto.NewArticleRequest) *storeNewArticleService {
	return &storeNewArticleService{
		AppContext: appCtx,
		ReqBody:    reqBody,
	}
}

func (s *storeNewArticleService) Run() error {
	article := models.Article{
		Author: s.ReqBody.Author,
		Title:  s.ReqBody.Title,
		Body:   s.ReqBody.Body,
	}
	if err := s.AppContext.DB.Create(&article).Error; err != nil {
		return err
	}

	ctx := context.Background()
	data, err := utils.Marshal(&article)
	if err != nil {
		return err
	}

	if err := s.AppContext.RedisClient.Publish(ctx, "article_created", data).Err(); err != nil {
		log.Error(&logrus.Fields{"error": err.Error()}, "Something error when publish article_created")
	}

	return nil
}
