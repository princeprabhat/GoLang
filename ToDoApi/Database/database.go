package database

import (
	"database/sql"
	"log"
)

func InitDb() *sql.DB {

	db, err := sql.Open("sqlite3", "./master.db")
	if err != nil {
		log.Fatal(err)
	}
	// Close the database connection

	// Create the table if it doesn't exist
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS todo (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		body TEXT
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s", err, sqlStmt)
		return nil
	}

	return db

}
