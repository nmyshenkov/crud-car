package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	DSN = "username:password@tcp(127.0.0.1:3306)/vehicals"
)

// InitDB - инициализация соединения с БД
func InitDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
