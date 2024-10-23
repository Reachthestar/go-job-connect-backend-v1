package users

import "time"

type User struct {
    ID                int64    `json:"id"`  
    FirstName         string   `json:"firstName"`
    LastName          string   `json:"lastName"`
    Email             string   `json:"email"`
    Password          string   `json:"password"`
    Phone             *string   `json:"phone"`
    Province          *string   `json:"province"`
    City              *string   `json:"city"`
    JobTitle          *string   `json:"jobTitle"`
    ProfileImage      *string   `json:"profileImage"`
    CompanyName       *string   `json:"companyName"`
    CompanyDescription *string  `json:"companyDescription"`
    IsActive          int       `json:"isActive"`
	CreatedAt		time.Time   `json:"createdAt"`
	UpdatedAt     	time.Time   `json:"updatedAt"`
    Role              string   `json:"role"`
}

// type User struct {
//     ID                int64   
//     FirstName         string   `binding:"required"`
//     LastName          string   `binding:"required"`
//     Email             string   `binding:"required"`
//     Password          string   `binding:"required"`
//     Phone             string   
//     Province          string   
//     City              string   
//     JobTitle          string   
//     ProfileImage      string   
//     CompanyName       string   
//     CompanyDescription string  
//     IsActive          int       
// 	CreatedAt		time.Time
// 	UpdatedAt     	time.Time
//     Role              string   `binding:"required"`
// }




