package controller

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

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

func (c *Controller) GarbageCollector() {
	ticker := time.NewTicker(time.Hour * 1)
	defer ticker.Stop()
	quit := make(chan bool)
	for {
		select {
		case <-ticker.C:
			c.RemoveDeadLinks()
		case <-quit:
			return
		}
	}
}

func (c *Controller) RemoveDeadLinks() {
	err := c.DB.DeleteExpiredLinks(context.TODO())
	if err != nil {
		log.Println("Failed to delete records from database")
		return
	}
	log.Println("Successfully deleted records from database")
}
