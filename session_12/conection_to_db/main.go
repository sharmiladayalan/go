package main

import (
	"database/sql"
	"fmt"

// import _ "github.com/mattn/go-sqlite3"
// → load sqlite driver (register only)

	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

var schema = `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    hashed_password BLOB NOT NULL, -- Storing as BLOB for byte slice
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
`

func main() {

	dbName := "data.db"

//Very important you do not do this in prod 
// → delete file, "_" ignore error
	_ = os.Remove(dbName)

	
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		// → print error + stop program immediately
		log.Fatal(err)
	}

	defer func() {
		fmt.Println("closing database connection")
		if err := db.Close(); err != nil {
			log.Printf("error closing database connection: %v", err)
		}
	}()

	// → check real DB connection works
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("database connection established")

	_, err = db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("table was created")

}