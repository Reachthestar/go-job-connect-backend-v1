package main

import (
	"github.com/Reachthestar/go-job-connect-backend/modules/routes"
	"github.com/Reachthestar/go-job-connect-backend/pkg/databases"
	"github.com/gin-gonic/gin"
)

func main() {
	databases.InitDB()
	server := gin.Default()
	routes.UserRouter(server)
	server.Run(":8080")
}

