package usecase

import "github.com/stobita/golang-ca-sample/domain"

type CommentInteractor struct {
	CommentRespository CommentRespository
}

func (interactor *CommentInteractor) Add(c domain.Comment) (err error) {
	_, err = interactor.CommentRespository.Store(c)
	return
}

func (interactor *CommentInteractor) Comments() (comment domain.Comments, err error) {
	comment, err = interactor.CommentRespository.FindAll()
	return
}

func (interactor *CommentInteractor) CommentById(identifier int) (comment domain.Comment, err error) {
	comment, err = interactor.CommentRespository.FindById(identifier)
	return
}
