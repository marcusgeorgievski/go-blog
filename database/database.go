// database/database.go

package database

import (
	"database/sql"
	. "goblog/logger"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() (*sql.DB, error) {
	connStr := "postgresql://marcusgeorgievski:6DErYHb4kwmA@ep-wispy-cherry-a5cakkha.us-east-2.aws.neon.tech/neondb?sslmode=require"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		Error.Println("Error opening database connection:", err)
		return nil, err
	}
	
	
	Info.Println("Connected to database: ", db.Stats().InUse)
	return db, nil
}

func GetDB() *sql.DB {
    return db
}

func SetDB(newDB *sql.DB) {
    db = newDB
}
