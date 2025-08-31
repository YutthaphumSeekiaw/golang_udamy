package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db") // fix here
	if err != nil {
		panic("Cannot Connect Database!")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUserTable := `
	  CREATE TABLE IF NOT EXISTS users (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    email TEXT NOT NULL UNIQUE,
	    password TEXT NOT NULL
	  )
	`
	_, err := DB.Exec(createUserTable)

	if err != nil {
		panic("Could not create table users !")
	}

	createEventTable := `
	  CREATE TABLE IF NOT EXISTS events (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    name TEXT NOT NULL,
	    description TEXT NOT NULL,
	    location TEXT NOT NULL,
	    dateTime DATETIME NOT NULL,
	    userId INTEGER,
		FOREIGN KEY(userId) REFERENCES users(id)
	  )
	`

	_, err = DB.Exec(createEventTable)

	if err != nil {
		panic("Could not create table events !")
	}

	createRegistrationTable := `
	  CREATE TABLE IF NOT EXISTS registrations (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    eventId INTEGER,
	    userId INTEGER,
		FOREIGN KEY(eventId) REFERENCES events(eventId),
		FOREIGN KEY(userId) REFERENCES users(userId)
	  )
	`

	_, err = DB.Exec(createRegistrationTable)

	if err != nil {
		panic("Could not create table registration table !")
	}
}
