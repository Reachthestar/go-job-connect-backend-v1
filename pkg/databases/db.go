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
    	email TEXT NOT NULL UNIQUE,
    	password TEXT NOT NULL,
    	phone TEXT,
    	province TEXT,
    	city TEXT,
    	jobTitle TEXT,
    	profileImage TEXT,
    	companyName TEXT UNIQUE,
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
}

