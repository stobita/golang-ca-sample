package infra

import (
	"github.com/gin-gonic/gin"
	"github.com/stobita/golang-ca-sample/interface/controller"
)

var Router *gin.Engine

func init() {
	router := gin.Default()
	articleController := controller.NewArticleController(NewSqlHandler())
	commentController := controller.NewCommentController(NewSqlHandler())
	router.POST("/articles", func(c *gin.Context) { articleController.Create(c) })
	router.GET("/articles", func(c *gin.Context) { articleController.Index(c) })
	router.GET("/articles/:id", func(c *gin.Context) { articleController.Show(c) })
	router.POST("/articles/:id/comments", func(c *gin.Context) { commentController.Create(c) })
	router.GET("comments", func(c *gin.Context) { commentController.Index(c) })
	router.GET("comments/:id", func(c *gin.Context) { commentController.Show(c) })
	Router = router
}
