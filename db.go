package main

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

func connectPool(connStr string) error {
	var err error
	DB, err = pgxpool.Connect(context.Background(), connStr)
	return err
}
