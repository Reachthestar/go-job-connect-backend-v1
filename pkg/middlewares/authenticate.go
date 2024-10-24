package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Reachthestar/go-job-connect-backend/modules/users"
	"github.com/Reachthestar/go-job-connect-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if len(token) > 7 && strings.HasPrefix(token, "Bearer ") {
		token = token[7:]
	} else {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}


	//Verify Token
	userId, err := utils.VerifyToken(token)

	if err != nil {
		fmt.Println("authorized error :", err)
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	// Call the GetUserByID 
	user,err := users.GetUserByID(userId)
	if err != nil {
		fmt.Println(err)
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Could not find this user id"})
		return
	}	

	// set user to gin context
	context.Set("user", user)
	context.Next()
}