package routers

import (
	"cicd/internal/middleware"
	v1 "cicd/internal/routers/api/v1"

	_ "github.com/go-programming-tour-book/blog-service/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())
	r.POST("/auth", v1.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	{
		tag := v1.NewTag()
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		article := v1.NewArticle()
		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/artivles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/stage", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}
	return r
}
