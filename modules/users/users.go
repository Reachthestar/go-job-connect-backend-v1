package users

import "time"

type User struct {
    ID                int64      
    FirstName         string   `binding:"required"`
    LastName          string   `binding:"required"`
    Email             string   `binding:"required"`
    Password          string   `binding:"required"`
    Phone             string   
    Province          string   
    City              string   
    JobTitle          string   
    ProfileImage      string   
    CompanyName       string   
    CompanyDescription string  
    IsActive          int       
	CreatedAt		time.Time
	UpdatedAt     	time.Time
    Role              string   `binding:"required"`
}




