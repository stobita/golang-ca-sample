package usecase

import "github.com/stobita/golang-ca-sample/domain"

type CommentRespository interface {
	Store(domain.Comment) (int, error)
	FindById(int) (domain.Comment, error)
	FindAll() (domain.Comments, error)
}
