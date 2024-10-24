package users

import (
	"fmt"
	"net/http"

	"github.com/Reachthestar/go-job-connect-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)


func Signup(context *gin.Context) {
	var user User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse request data."})
		return
	}

	err = user.Save()

	if err != nil {
		fmt.Println("Could not save user",err)
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not save user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message":"User created successfully","user":user})
}

func Login(context *gin.Context) {
	var user User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse request data."})
		return
	}

	// validate User
	err = user.ValidateCredentials()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusUnauthorized, gin.H{"message":"Could not authenticate user."})
		return
	}

	//Generate Token
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
	context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not authenticate user."})
	return
	}

	context.JSON(http.StatusOK, gin.H{"message":"Login successfully", "token":token})	
}


