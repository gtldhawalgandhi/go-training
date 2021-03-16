package db

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

// PGStore ...
type PGStore struct {
	db *pgxpool.Pool
}

// NewPGStore creates PGStore that implements the 'Store' interface
func NewPGStore(db *pgxpool.Pool) Store {
	return &PGStore{
		db: db,
	}
}
