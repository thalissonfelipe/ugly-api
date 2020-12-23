# Ugly REST API developed with golang

This is just a simple rest api developed with go for learning purposes.

## Endpoints

* Get all movies
  `GET /api/v1/movies`

* Get a movie by ID
  `GET /api/v1/movies/{id}`

* Create a new movie
  `POST /api/v1/movies`

* Update a movie by ID
  `PUT /api/v1/movies/{id}`

* Delete a movie by ID
  `DELETE /api/v1/movies/{id}`

## How to run

- Clone this repository
- Run `go mod download`
- Run `go build`
- Run `./ugly-api`

## TODO

- [x] Logging
- [ ] Add Request ID
- [ ] Body validation
- [ ] Data persistence (MongoDB)
- [ ] Handle Errors
- [ ] Set environment variables
- [ ] Set up a project architecture
- [ ] Unit Tests