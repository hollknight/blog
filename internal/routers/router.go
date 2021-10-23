package routers

import (
	"blog/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	article := v1.NewArticle()
	tag := v1.NewTag()

	apiv1 := r.Group("/api/v1")
	{
		tagsGroup := apiv1.Group("/tags")
		{
			tagsGroup.POST("", tag.Create)
			tagsGroup.DELETE("/:id", tag.Delete)
			tagsGroup.PUT("/:id", tag.Update)
			tagsGroup.PATCH("/:id/state", tag.Update)
			tagsGroup.GET("", tag.List)
		}
		articlesGroup := apiv1.Group("/articles")
		{
			articlesGroup.POST("", article.Create)
			articlesGroup.DELETE("/:id", article.Delete)
			articlesGroup.PUT("/:id", article.Update)
			articlesGroup.PATCH("/:id/state", article.Update)
			articlesGroup.GET("/:id", article.Get)
			articlesGroup.GET("", article.List)
		}

	}

	return r
}
