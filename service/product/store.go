package product

import "database/sql"

type Store struct {
	db *sql.DB
}

// Repository for the products
func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}