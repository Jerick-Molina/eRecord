package db

import "database/sql"

type Record struct {
	db *sql.DB
	*Queries
}

func NewRecord(db *sql.DB) *Record {
	return &Record{
		db:      db,
		Queries: New(db),
	}
}
