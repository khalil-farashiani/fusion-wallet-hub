package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// SeedData holds the data to seed the database.
var SeedData = []struct {
	AccountID string
	Amount    int
}{
	{"123", 1000},
	{"1234", 1500},
	{"1235", 1200},
	{"1236", 1000},
	{"1237", 1500},
	{"1238", 1200},
	{"1239", 1000},
	{"1240", 1500},
	{"1241", 1200},
	{"1242", 1000},
	{"1243", 1500},
	{"1244", 1200},
	{"1245", 1000},
	{"1246", 1500},
	{"1247", 1200},
	{"1248", 1000},
	{"1249", 1500},
	{"1250", 1200},
}

// seedDatabase seeds the balance table with initial data.
func seedDatabase(db *sql.DB) error {
	query := `INSERT INTO balance (account_id, amount, created_at, updated_at) VALUES (?, ?, ?, ?)`

	for _, data := range SeedData {
		_, err := db.Exec(query, data.AccountID, data.Amount, time.Now(), time.Now())
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	//TODO: get this from config
	dsn := "testuser:veryhardhardpasswordformysql@tcp(localhost:3307)/wallet_db"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// Seed the database
	if err := seedDatabase(db); err != nil {
		log.Fatalf("Error seeding database: %v", err)
	} else {
		log.Println("Database seeded successfully!")
	}
}
