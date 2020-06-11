package server

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) initRoutes(server *gin.Engine) {

	v1 := server.Group("/api/v1")
	{
		v1.GET("/status", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})
	}

	extractorGroup := v1.Group("/extractor")
	{
		testGroup := extractorGroup.Group("/test")
		{
			testGroup.POST("", s.Controllers.Test.Create)
			testGroup.GET("", s.Controllers.Test.Get)
			testGroup.GET("/:id", s.Controllers.Test.GetById)
			testGroup.PUT("/:id", s.Controllers.Test.Update)
			testGroup.DELETE("/:id", s.Controllers.Test.Delete)
		}
	}
}
