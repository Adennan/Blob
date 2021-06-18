package router

import (
	_ "Blob/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	url := ginSwagger.URL("http://127.0.0.1:8000/swagger/docs.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	v1 := r.Group("/v1")
	{
		// tags
		v1.POST("/tags")
		v1.DELETE("/tags/:id")
		v1.GET("/tags")
		v1.GET("/tag/status")

		// article
		v1.POST("/articles")
		v1.POST("/articles/:id")
		v1.DELETE("/articles/:id")
		v1.GET("/articles/:id")
		v1.GET("/articles")
	}

	return r
}
