package routers

import (
	_ "blog/docs"
	"blog/internal/middleware"
	"blog/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())

	url := ginSwagger.URL("http://127.0.0.1:8200/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

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
