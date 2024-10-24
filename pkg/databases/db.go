package databases

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)
var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3","pkg/databases/go-job-connect-v1.db")

		if err != nil {
		panic("Could not connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
    	firstName TEXT NOT NULL,
    	lastName TEXT NOT NULL,
    	email TEXT NOT NULL UNIQUE CHECK(length(email) <= 50),
    	password TEXT NOT NULL CHECK(length(password) <= 256),
    	phone TEXT CHECK(length(phone) <= 20),
    	province TEXT CHECK(length(city) <= 50),
    	city TEXT CHECK(length(city) <= 50),
    	jobTitle TEXT CHECK(length(jobTitle) <= 100),
    	profileImage TEXT CHECK(length(profileImage) <= 256),
    	companyName TEXT UNIQUE CHECK(length(companyName) <= 100),
    	companyDescription TEXT,
    	isActive INTEGER NOT NULL DEFAULT 1,
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    	role TEXT CHECK( role IN ('SEEKER', 'COMPANY') )  NOT NULL
		)
		`
		_, err := DB.Exec(createUsersTable)
		if err != nil {
			panic("Could not create users table.")
		}

	createJobsTable := `
	CREATE TABLE IF NOT EXISTS jobs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
    	description TEXT NOT NULL,
    	qualification TEXT NOT NULL,
    	title TEXT NOT NULL CHECK(length(title) <= 100),
    	position TEXT CHECK(length(position) <= 100),
    	type TEXT NOT NULL CHECK(type IN ('FULL_TIME', 'PART_TIME')),
    	province TEXT CHECK(length(province) <= 50),
    	city TEXT CHECK(length(city) <= 50),
    	salaryMin TEXT CHECK(length(salaryMin) <= 10),
    	salaryMax TEXT CHECK(length(salaryMax) <= 10),
    	createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    	updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    	userId INTEGER,
    	FOREIGN KEY (userId) REFERENCES users(id)
		)`

		_, err = DB.Exec(createJobsTable)
		if err != nil {
			panic("Could not create jobs table.")
		}

	createApplicationsTable := `
	CREATE TABLE IF NOT EXISTS applications (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
    	status TEXT NOT NULL CHECK(status IN ('PENDING', 'ACCEPTED', 'DENIED')),
    	createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    	updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    	userId INTEGER NOT NULL,
    	jobId INTEGER NOT NULL,
    	FOREIGN KEY (userId) REFERENCES users(id),
    	FOREIGN KEY (jobId) REFERENCES jobs(id)
		)`

		_, err = DB.Exec(createApplicationsTable)
		if err != nil {
			panic("Could not create jobs table.")
		}

	createExperiencesTable := `
	CREATE TABLE IF NOT EXISTS experiences (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
    	title TEXT NOT NULL CHECK(title IN ('EDUCATION', 'WORK')),
    	name TEXT NOT NULL CHECK(length(name) <= 100),
    	description TEXT,
    	startYear INTEGER,
    	endYear INTEGER,
    	userId INTEGER,
    	FOREIGN KEY (userId) REFERENCES users(id)
		)`

		_, err = DB.Exec(createExperiencesTable)
		if err != nil {
			panic("Could not create jobs table.")
		}
}

