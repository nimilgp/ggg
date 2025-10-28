package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Post interface {
		Create(context.Context) error
	}
	Users interface {
		Create(context.Context) error
	}
}

func NewPostgresStorage(db *sql.DB) Storage {
	return Storage{
		Post:  &PostgresPostStore{db: db},
		Users: &PostgresUserStore{db: db},
	}
}
