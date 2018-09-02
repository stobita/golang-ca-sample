package usecase

import "github.com/stobita/golang-ca-sample/domain"

type ArticleInteractor struct {
	ArticleRepository ArticleRepository
}

func (interactor *ArticleInteractor) Add(a domain.Article) (err error) {
	_, err = interactor.ArticleRepository.Store(a)
	return
}

func (interactor *ArticleInteractor) Articles() (article domain.Articles, err error) {
	article, err = interactor.ArticleRepository.FindAll()
	return
}

func (interactor *ArticleInteractor) ArticleById(identifier int) (article domain.Article, err error) {
	article, err = interactor.ArticleRepository.FindById(identifier)
	return
}
