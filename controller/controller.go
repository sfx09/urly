package controller

import (
	"database/sql"
	"errors"
	"net/http"

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

func (c *Controller) HandleCreateLink(w http.ResponseWriter, r *http.Request) {}

func (c *Controller) HandleQueryLink(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) HandleRedirectLink(w http.ResponseWriter, r *http.Request) {

}
