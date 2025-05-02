package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDatabase(path string) (*sql.DB, error) {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return nil, fmt.Errorf("failed to create database directory: %w", err)
	}
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil
}

func InitDatabase(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS Dishes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		img TEXT,
		calorie INTEGER NOT NULL,
		runtime INTEGER NOT NULL,
		like REAL DEFAULT 0,
		dislike REAL DEFAULT 0,
		nationality TEXT
	);
	`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create dishes table: %w", err)
	}
	return nil
}