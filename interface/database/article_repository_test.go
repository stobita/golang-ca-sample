package database

import (
	"testing"

	"github.com/stobita/golang-ca-sample/domain"
)

func TestStore(t *testing.T) {
	article := new(domain.Article)
	repo := new(ArticleRepository)
	result, err := repo.Store(*article)
	if err != nil {
		t.Log(err)
		t.Error("Store error")
	}
	if result == 0 {
		t.Error("Store error")
	}
}
