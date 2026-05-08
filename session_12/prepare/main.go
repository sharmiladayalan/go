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

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword string
	CreatedAt      time.Time
}

func main() {

	dbName := "users_database.db"

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ database connection established")

	// 👉 Create table (needed)
	createTable(db)

	// 👉 Context (kept as in your code)
	ctx := context.Background()

	// ❗ Using PREPARED version for insert (as you asked)
	userID, err := createUserWithPrepared(db, "Joseph Abah 2", "jo2@localhost.com", "test09120")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("✅ User inserted with ID:", userID)

	// 👉 Fetch user (normal function, no prepare)
	user, err := GetUserByEmail(db, "jo2@localhost.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("✅ User fetched: %+v\n", user)

	// 👉 Fetch all users
	users, err := GetUsers(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ All users:")
	for _, u := range users {
		fmt.Printf("%+v\n", u)
	}

	// 👉 Just to show ctx version exists (not used for insert now)
	_ = ctx
}

// 👉 Create table
func createTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT UNIQUE,
		hashed_password TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

// 👉 Your original ctx version (kept unchanged)
func createUserWithCtx(ctx context.Context, db *sql.DB, name, email, hashedPassword string) (int64, error) {
	stmt, err := db.Prepare(`INSERT INTO users (name, email, hashed_password) VALUES (?, ?, ?)`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	hp, err := bcrypt.GenerateFromPassword([]byte(hashedPassword), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	result, err := stmt.ExecContext(ctx, name, email, string(hp))
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// 👉 Using THIS for insert (prepared)
func createUserWithPrepared(db *sql.DB, name, email, hashedPassword string) (int64, error) {
	stmt, err := db.Prepare(`INSERT INTO users (name, email, hashed_password) VALUES (?, ?, ?)`)
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
	return result.LastInsertId()
}

// 👉 Normal select (no prepare, as you wanted)
func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	stmt := `SELECT id, name, email, hashed_password, created_at FROM users WHERE email = ?`

	row := db.QueryRow(stmt, email)

	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.HashedPassword, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUsers(db *sql.DB) ([]User, error) {
	stmt := `SELECT id, name, email, hashed_password, created_at FROM users`
	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.HashedPassword,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}


