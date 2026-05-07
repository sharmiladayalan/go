package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

// 👉 USERS TABLE (missing in your code)
var userSchema = `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    email TEXT UNIQUE,
    hashed_password TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
`

// 👉 PROFILE TABLE (fixed name + FK)
var profileSchema = `
CREATE TABLE IF NOT EXISTS profile (
    user_id INTEGER PRIMARY KEY,
    avatar TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);
`

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword string
	CreatedAt      time.Time
	Profile        Profile
}

type Profile struct {
	UserID  int
	Avatar  string
	Created time.Time
}

func main() {

	dbName := "users_database.db"

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ database connection established")

	// 👉 CREATE TABLES (you missed this)
	createTables(db)

	// 👉 Create user with transaction
	userID, err := createUser(db, "test with defer", "test@test.com", "password123", "avatar.png")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("✅ User created with ID:", userID)

	// 👉 Fetch user
	user, err := GetUserByEmail(db, "test@test.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("✅ User fetched: %+v\n", user)
}

// 👉 Create tables
func createTables(db *sql.DB) {
	_, err := db.Exec(userSchema)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(profileSchema)
	if err != nil {
		log.Fatal(err)
	}
}

// 👉 TRANSACTION FUNCTION (your logic preserved)
func createUser(db *sql.DB, name, email, hashedPassword, avatar string) (int64, error) {
	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, `INSERT INTO users (name, email, hashed_password) VALUES (?, ?, ?)`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	hp, err := bcrypt.GenerateFromPassword([]byte(hashedPassword), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(name, email, string(hp))
	if err != nil {
		return 0, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// 👉 FIXED: table name = profile (not profiles)
	profileStmt, err := tx.PrepareContext(ctx, `INSERT INTO profile (user_id, avatar) VALUES (?, ?)`)
	if err != nil {
		return 0, err
	}
	defer profileStmt.Close()

	_, err = profileStmt.Exec(userID, avatar)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return userID, nil
}

// 👉 JOIN QUERY (fixed table name)
func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	stmt := `SELECT u.id, u.name, u.email, u.hashed_password, u.created_at, p.avatar
	FROM users u
	INNER JOIN profile p ON u.id = p.user_id
	WHERE u.email = ?`

	row := db.QueryRow(stmt, email)

	var user User
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.HashedPassword,
		&user.CreatedAt,
		&user.Profile.Avatar,
	)
	if err != nil {
		return nil, err
	}

	user.Profile.UserID = user.ID

	return &user, nil
}
