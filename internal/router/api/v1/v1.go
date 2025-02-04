package v1

import (
	"github.com/onepiece010938/go-line-message-analyzer/internal/app"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.RouterGroup, app *app.Application) {
	v1 := router.Group("/v1")
	{
		v1.POST("/callback", Callback(app))
		v1.GET("/sample", SAMPLE)
		v1.GET("/analyze", StartAnalyze(app))
		// chatGPT
		v1.POST("/chatGPT", ChatGPT(app))
	}
}
