package controller

import (
	"strconv"

	"github.com/stobita/golang-ca-sample/domain"
	"github.com/stobita/golang-ca-sample/interface/database"
	"github.com/stobita/golang-ca-sample/usecase"
)

type CommentController struct {
	Interactor usecase.CommentInteractor
}

func NewCommentController(sqlHandler database.SqlHandler) *CommentController {
	return &CommentController{
		Interactor: usecase.CommentInteractor{
			CommentRespository: &database.CommentRespository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *CommentController) Create(c Context) {
	articleId, _ := strconv.Atoi(c.Param("id"))
	cm := domain.Comment{}
	c.Bind(&cm)
	cm.ArticleID = articleId
	err := controller.Interactor.Add(cm)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, nil)
}

func (controller *CommentController) Index(c Context) {
	comments, err := controller.Interactor.Comments()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, comments)
}

func (controller *CommentController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	comment, err := controller.Interactor.CommentById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, comment)
}
