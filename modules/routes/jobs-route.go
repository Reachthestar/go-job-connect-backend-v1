package routes

import (
	"github.com/Reachthestar/go-job-connect-backend/modules/jobs"
	"github.com/Reachthestar/go-job-connect-backend/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func JobRouter(server *gin.Engine) {
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/jobs/createJob", jobs.CreateJob)
	authenticated.DELETE("/jobs/:id", )
}