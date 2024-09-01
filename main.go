package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/sfx09/urly/controller"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")
	conn := os.Getenv("CONN")

	c, err := controller.NewController(conn)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	server := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	mux.HandleFunc("POST /create", c.HandleCreateLink)
	mux.HandleFunc("GET /query/{id}", c.HandleQueryLink)
	mux.HandleFunc("GET /{id}", c.HandleRedirectLink)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
