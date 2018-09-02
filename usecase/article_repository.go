package usecase

import "github.com/stobita/golang-ca-sample/domain"

type ArticleRepository interface {
	Store(domain.Article) (int, error)
	FindById(int) (domain.Article, error)
	FindAll() (domain.Articles, error)
}
