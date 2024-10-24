package users

import (
	"errors"
	"fmt"

	"github.com/Reachthestar/go-job-connect-backend/pkg/databases"
	"github.com/Reachthestar/go-job-connect-backend/pkg/utils"
)


func (u *User) Save() error{
	// Validate user data
	if u.FirstName == "" {
		return errors.New("First name is required")
	}
	if u.LastName == "" {
		return errors.New("Last name is required")
	}
	if u.Email == "" {
		return errors.New("Email is required")
	}
	if u.Password == "" {
		return errors.New("Password is required")
	}
	if u.Role == "" {
		return errors.New("Role is required")
	}

	query := "INSERT INTO users(firstName, lastName, email, password, role) VALUES (?, ?, ?, ?, ?)"
	stmt, err := databases.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	//Hash password
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.FirstName, u.LastName, u.Email, hashedPassword, u.Role)

	if err != nil {
		fmt.Println("Error executing query:", err)
		return err
	}

	userId, err := result.LastInsertId() // return int64

	u.ID = userId

	return err
}

func (u *User) ValidateCredentials() error{
	// Validate email or password
	if u.Email == "" || u.Password == "" {
        return errors.New("email and password are required")
    }

	query := "SELECT id, password FROM users WHERE email = ?"
	row := databases.DB.QueryRow(query, u.Email)

	var retrievedPassword string

	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return errors.New("Credentials invalid")
	}

	//Compare password
	isPasswordMatch := utils.CompareHashPassword(u.Password, retrievedPassword)
	
	if !isPasswordMatch {
		return errors.New("Credentials invalid not match")
	}

	return nil
}

func GetUserByID(userId int64) (*User, error){
	var user User

	query := "SELECT id, firstName, lastName, email, phone, province, city, jobTitle, profileImage, companyName, companyDescription, isActive, createdAt, updatedAt, role FROM users WHERE id = ?"

	row := databases.DB.QueryRow(query, userId)
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Phone, &user.Province, &user.City, &user.JobTitle, &user.ProfileImage, &user.CompanyName, &user.CompanyDescription, &user.IsActive, &user.CreatedAt, &user.UpdatedAt, &user.Role)
	if err != nil {
	return nil,err 
}

return &user, nil
}

