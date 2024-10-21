package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func Signup(context *gin.Context) {
	var user User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse request data."})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not save user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message":"User created successfully"})
}

// func GetMe(context *gin.Context){

// 	context.JSON(http.StatusOK, gin.H{"message":"it is working!"})
// }