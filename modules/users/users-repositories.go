package users

import (
	"github.com/Reachthestar/go-job-connect-backend/pkg/databases"
	"github.com/Reachthestar/go-job-connect-backend/pkg/utils"
)


func (u User) Save() error{
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
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId

	return err
}