package main

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lemon-mint/pkgserver/db"
)

var DB *pgxpool.Pool
var DBQueries *db.Queries

func connectPool(connStr string) error {
	var err error
	DB, err = pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return err
	}
	DBQueries = db.New(DB)
	return nil
}
