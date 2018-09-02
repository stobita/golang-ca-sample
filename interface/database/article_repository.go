package database

import (
	"github.com/stobita/golang-ca-sample/domain"
)

type ArticleRepository struct {
	SqlHandler
}

func (repo *ArticleRepository) Store(a domain.Article) (id int, err error) {
	result, err := repo.Execute(
		"INSERT INTO article (title, content) VALUE (?,?)", a.Title, a.Content,
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

func (repo *ArticleRepository) FindById(identifier int) (article domain.Article, err error) {
	row, err := repo.Query("SELECT article.id, title, content, comment.id, comment.text FROM article INNER JOIN comment ON comment.article_id = article.id WHERE article.id = ?", identifier)
	defer row.Close()
	if err != nil {
		return
	}
	var id int
	var title string
	var content string
	var comments domain.Comments
	var commentId int
	var commentText string
	row.Next()
	if err = row.Scan(&id, &title, &content, &commentId, &commentText); err != nil {
		return
	}
	comment := domain.Comment{
		ID:   commentId,
		Text: commentText,
	}
	comments = append(comments, comment)
	for row.Next() {
		if err := row.Scan(nil, nil, nil, &commentId, &commentText); err != nil {
			continue
		}
		comment := domain.Comment{
			ID:   commentId,
			Text: commentText,
		}
		comments = append(comments, comment)
	}
	article.ID = id
	article.Title = title
	article.Content = content
	article.Comments = comments
	return
}

func (repo *ArticleRepository) FindAll() (articles domain.Articles, err error) {
	rows, err := repo.Query("SELECT id, title, content FROM article")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int
		var title string
		var content string
		if err := rows.Scan(&id, &title, &content); err != nil {
			continue
		}
		article := domain.Article{
			ID:      id,
			Title:   title,
			Content: content,
		}
		articles = append(articles, article)

	}
	return
}
