package jobs

import (
	"errors"
	"fmt"

	"github.com/Reachthestar/go-job-connect-backend/pkg/databases"
)

var jobs = []Job{}

func (j *Job)Save() error{

	if j.Title == "" {
		return errors.New("Job title is required.")		
	}
	if j.Description == "" {
		return errors.New("Job description is required.")		
	}
	if j.Qualification == "" {
		return errors.New("Job qualification is required.")		
	}
	if j.Type == "" {
		return errors.New("Job type is required.")		
	}
	if j.SalaryMin == nil || j.SalaryMax == nil {
		return errors.New("Both minimum and maximum salary are required.")
	}

	query := `
	INSERT INTO jobs(description, qualification, title, position, type, province, city, salaryMin, salaryMax, userId)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	stmt, err := databases.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(j.Description, j.Qualification, j.Title, j.Position, j.Type, j.Province, j.City, j.SalaryMin, j.SalaryMax, j.UserID)
	if err != nil {
		return err
	}

	// get ID
	id, err := result.LastInsertId()
	j.ID = id

	fmt.Println(j.ID)
	return err
}