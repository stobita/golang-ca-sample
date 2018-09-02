package database

import (
	"github.com/stobita/golang-ca-sample/domain"
)

type CommentRespository struct {
	SqlHandler
}

func (repo *CommentRespository) Store(c domain.Comment) (id int, err error) {
	result, err := repo.Execute(
		"INSERT INTO comment (content, article_id) VALUE (?,?)", c.Text, c.ArticleID,
	)
	if err != nil {
		return
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	return
}

func (repo *CommentRespository) FindById(identifier int) (comment domain.Comment, err error) {
	row, err := repo.Query("SELECT id, text, article_id FROM comment WHERE id = ?", identifier)
	defer row.Close()
	if err != nil {
		return
	}
	var id int
	var text string
	var articleID int
	row.Next()
	if err = row.Scan(&id, &text, &articleID); err != nil {
		return
	}
	comment.ID = id
	comment.Text = text
	comment.ArticleID = articleID
	return
}

func (repo *CommentRespository) FindAll() (comments domain.Comments, err error) {
	rows, err := repo.Query("SELECT id, text, article_id FROM comment")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int
		var text string
		var articleID int
		if err = rows.Scan(&id, &text, &articleID); err != nil {
			return
		}
		comment := domain.Comment{
			ID:        id,
			Text:      text,
			ArticleID: articleID,
		}
		comments = append(comments, comment)
	}
	return
}
