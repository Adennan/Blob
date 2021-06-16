package router

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

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
