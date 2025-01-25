package routes

import (
	"open-highscore/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/scores", controllers.GetScores)
	r.POST("/scores", controllers.AddScore)
}
