package storage

import (
	"fmt"
	"github.com/hararudoka/gotemplate/internal/env"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	*sqlx.DB
}

// Open opens database connection
func Open(e *env.Env) (*DB, error) {
	conn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		e.Username, e.Password, e.Hostname, e.DBName, e.Mode)

	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		return nil, err
	}

	return &DB{DB: db}, nil
}
