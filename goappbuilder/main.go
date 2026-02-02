package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func handler(w http.ResponseWriter, r *http.Request) {
	// Increment counter
	_, err := db.Exec("UPDATE counter SET visits = visits + 1 WHERE id = 1")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Read current counter
	var visits int
	err = db.QueryRow("SELECT visits FROM counter WHERE id = 1").Scan(&visits)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, "Page visits: %d\n", visits)
}

func main() {
	// Read database credentials from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	if dbUser == "" || dbPassword == "" || dbHost == "" || dbName == "" {
		log.Fatal("Database environment variables (DB_USER, DB_PASSWORD, DB_HOST, DB_NAME) are required")
	}

	// Build DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPassword, dbHost, dbName)

	// Retry connecting to MySQL until it's ready
	var err error
	for {
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Printf("Error opening database: %v. Retrying in 3s...", err)
			time.Sleep(3 * time.Second)
			continue
		}

		err = db.Ping()
		if err != nil {
			log.Printf("Cannot connect to MySQL: %v. Retrying in 3s...", err)
			time.Sleep(3 * time.Second)
			continue
		}

		break
	}

	log.Println("Connected to MySQL successfully!")

	// Start HTTP server
	http.HandleFunc("/", handler)
	log.Println("Listening on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

