package controller

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/sfx09/urly/database"
	"github.com/sfx09/urly/internal"
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

func (c *Controller) HandleCreateLink(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Url string `json:"url"`
	}
	defer r.Body.Close()
	req := Request{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Failed to decode request JSON")
		return
	}
	if !internal.IsValidUrl(req.Url) {
		respondWithError(w, http.StatusBadRequest, "Invalid URL")
		return
	}

	shortLink := internal.GenerateRandomString()
	log.Println(shortLink, req.Url)
	link, err := c.DB.CreateLink(r.Context(), database.CreateLinkParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		ShortLink: shortLink,
		FullLink:  req.Url,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to write to DB")
		return
	}
	respondWithJSON(w, http.StatusOK, link)
}

func (c *Controller) HandleQueryLink(w http.ResponseWriter, r *http.Request) {
	shortLink := r.PathValue("id")
	if shortLink == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid request URL")
		return
	}
	link, err := c.DB.GetByShortLink(r.Context(), shortLink)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Failed to find short link")
		return
	}
	respondWithJSON(w, http.StatusOK, link)
}

func (c *Controller) HandleRedirectLink(w http.ResponseWriter, r *http.Request) {
	shortLink := r.PathValue("id")
	if shortLink == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid request URL")
		return
	}
	link, err := c.DB.GetByShortLink(r.Context(), shortLink)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Failed to find short link")
		return
	}
	http.Redirect(w, r, link.FullLink, http.StatusTemporaryRedirect)
}
