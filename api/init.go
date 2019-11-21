package api

import "database/sql"

// Typ - главная структура
type Typ struct {
	db *sql.DB
}

// Init - инициализация главной структуры
func Init(db *sql.DB) Typ {
	return Typ{
		db: db,
	}
}
