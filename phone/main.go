package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "nathan"
	dbName = "gophercises_phone"
)

func main() {
	password := os.Getenv("POSTGRES_PASS")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
		host, port, user, password)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Failed to open db: %v", err)
	}
	if err = resetDb(db, dbName); err != nil {
		log.Fatalf("Failed to reset db: %v", err)
	}
	db.Close()

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbName)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Failed to open db: %v", err)
	}
	defer db.Close()

	err = createPhoneTable(db)
	if err != nil {
		log.Fatalf("Failed to Create table db: %v", err)
	}

	id, err := insertPhone(db, "44555545445")
	if err != nil {
		log.Fatalf("Failed to insert phone number: %v", err)
	}

	fmt.Println(id)
}

// Database setup details
func createDb(db *sql.DB, name string) error {
	_, err := db.Exec("CREATE DATABASE " + name)
	if err != nil {
		return err
	}
	return nil
}

func resetDb(db *sql.DB, name string) error {
	_, err := db.Exec("DROP DATABASE IF EXISTS " + name)
	if err != nil {
		return err
	}
	return createDb(db, name)
}

func createPhoneTable(db *sql.DB) error {
	statement := `
		CREATE TABLE IF NOT EXISTS phone_numbers (
			id SERIAL,
			value VARCHAR(255)
		)
	`

	_, err := db.Exec(statement)
	return err
}

// Start phone specific database stuff
func insertPhone(db *sql.DB, phone string) (int, error) {
	statement := `INSERT INTO phone_numbers(value) VALUES($1) RETURNING id`
	var id int
	err := db.QueryRow(statement, phone).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

// Phone specific
func normalize(phone string) string {
	var buf bytes.Buffer
	for _, r := range phone {
		if r >= '0' && r <= '9' {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}
