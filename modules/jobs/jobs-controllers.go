package jobs

import (
	"fmt"
	"net/http"
	"strconv"

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

	context.JSON(http.StatusCreated, gin.H{"message": "Job created!", "job":job})
}

func DeleteJob(context *gin.Context) {
	jobId, err := strconv.ParseInt(context.Param("id"), 10, 64)	
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse job id."})
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

	userId := user.ID

	job, err := GetJobByID(jobId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Not found this job ID."})
		return
	}

	// check permission for the same userId
	if job.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message":"not authorized to delete job"})
		return
	}

	err = job.DeleteJob()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not delete job."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Job deleted successfully!"})
}