package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sfx09/urly/database"
	"github.com/sfx09/urly/internal"
)

func (c *Controller) HandleLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next(w, r)
	}
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
	// TODO: Validate short link exists
	err := c.DB.UpdateLinkCounter(r.Context(), database.UpdateLinkCounterParams{
		ShortLink: shortLink,
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to update database")
		return
	}
	link, err := c.DB.GetByShortLink(r.Context(), shortLink)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Failed to find short link")
		return
	}
	http.Redirect(w, r, link.FullLink, http.StatusTemporaryRedirect)
}
