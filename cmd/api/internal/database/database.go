package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewDB() (*sql.DB, error) {
	connStr := "postgresql://postgres:1234@localhost:5432/postgres?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	for _, query := range DBqueries {
		statement, err := db.Prepare(query)
		if err != nil {
			return nil, err
		}
		statement.Exec()
	}

	return db, nil
}
