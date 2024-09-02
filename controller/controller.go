package controller

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
	"github.com/sfx09/urly/database"
)

type Controller struct {
	DB *database.Queries
}

func NewController(conn string) (*Controller, error) {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return &Controller{}, errors.New("Failed to connect to database")
	}
	dbQueries := database.New(db)
	return &Controller{
		DB: dbQueries,
	}, nil
}
