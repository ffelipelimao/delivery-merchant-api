package db

import (
	"database/sql"

	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/config"
	_ "github.com/lib/pq"
)

func NewDatabase(config *config.Config) (*sql.DB, error) {
	connectionString := config.DBUrl

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
