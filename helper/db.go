// helper/helper.go
package helper

import (
	"log"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func ConnectDatabase() (*sqlx.DB, error) {
	connStr := "host=localhost user=postgres password=coder010203 dbname=task port=5432 sslmode=disable TimeZone=Asia/Seoul"
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	DB = db

	// Run DDL script for creating tables
	err = RunDDLScript()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func RunDDLScript() error {
	script := `
		CREATE TABLE IF NOT EXISTS blogs (
			id UUID PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			body TEXT NOT NULL
		);
		
		CREATE TABLE IF NOT EXISTS news (
			id UUID PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			body TEXT NOT NULL
		);
	`

	_, err := DB.Exec(script)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
