package service

import (
	"article-web-service/read-service/models"
	"article-web-service/shared/domains"
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/olivere/elastic/v7"
)

type listArticleService struct {
	AppContext *domains.Context
	Query      string
	Author     string
}

func NewListArticleService(appCtx *domains.Context, query string, author string) *listArticleService {
	return &listArticleService{
		AppContext: appCtx,
		Query:      query,
		Author:     author,
	}
}

func (s *listArticleService) Run() ([]models.Article, error) {
	var articles []models.Article
	query := elastic.NewBoolQuery()

	if len(s.Query) > 0 {
		query = query.Filter(elastic.NewMultiMatchQuery(s.Query, "title", "body").Type("phrase_prefix"))
	}

	if len(s.Author) > 0 {
		query = query.Filter(elastic.NewMatchPhrasePrefixQuery("author", s.Author))
	}

	src, err := query.Source()
	if err != nil {
		return make([]models.Article, 0), err
	}

	data, err := json.MarshalIndent(src, "", " ")
	if err != nil {
		return make([]models.Article, 0), err
	}

	fmt.Println(string(data))

	searchResult, err := s.AppContext.ESClient.Search().
		Index("articles").
		Query(query).
		Sort("created", false).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		return make([]models.Article, 0), err
	}

	var a models.Article
	for _, item := range searchResult.Each(reflect.TypeOf(a)) {
		article := item.(models.Article)
		articles = append(articles, article)
	}

	if len(articles) == 0 {
		return make([]models.Article, 0), nil
	}

	return articles, nil
}
