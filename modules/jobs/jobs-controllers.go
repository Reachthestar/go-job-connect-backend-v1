package jobs

import (
	"fmt"
	"net/http"

	"github.com/Reachthestar/go-job-connect-backend/modules/users"
	"github.com/gin-gonic/gin"
)

func CreateJob(context *gin.Context) {
	var job Job

	err := context.ShouldBindJSON(&job)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse request data."})
		return
	}

	// Retrieve the user from the context (set by Authenticate middleware)
	userInterface, ok := context.Get("user")
	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User not authenticated"})
		return
	}

	// Cast userInterface to *User struct
	user, ok := userInterface.(*users.User)
    if !ok {
        context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cast user object"})
        return
    }

	role := user.Role

	//Check role before allowing job creation
	if role != "COMPANY" {
		context.JSON(http.StatusForbidden, gin.H{"message": "Only COMPANY users can create jobs."})
		return
	}

	// Assign the userId to the job's UserID
	job.UserID = user.ID

	// Save the job
	err = job.Save()

	if err != nil {
		fmt.Println("Error saving job:", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not create jobs. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Job created!", "event":job})
}

