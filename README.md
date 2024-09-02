# URLY
A web service that converts long URLs into shorter, manageable links.

## FEATURES

* Create short manageable links for any URLs
* Query usage count for any short link
* Automatic deletion of unused links from database 

### USAGE 

* We use Goose for automatic database migrations

`go install github.com/pressly/goose/v3/cmd/goose@latest`

* We use sqlc for generating sql code in Go

`go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`


Take a look at the justfile to get started. 
