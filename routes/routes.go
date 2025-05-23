package routes

import (
	"github.com/gin-gonic/gin"
	"web-analyzer/handler"
)

func SetupRoutes(r *gin.Engine) {
	r.LoadHTMLGlob("web/*")
	r.GET("/", handler.LoadInitialPage)
	r.POST("/analyze", handler.InvokeAnalyzer)
}
