package routes

import (
	"github.com/Reachthestar/go-job-connect-backend/modules/users"
	"github.com/gin-gonic/gin"
)

func UserRouter(server *gin.Engine) {
	server.POST("/signup", users.Signup)
	// server.GET("/me",users.GetMe)
}