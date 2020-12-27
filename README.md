# Ugly REST API developed with golang

This is just a simple rest api developed with go for learning purposes.

## Endpoints

* Get all movies
  `GET /api/v1/movies`

* Get a movie by name
  `GET /api/v1/movies/{name}`

* Create a new movie
  `POST /api/v1/movies`

* Update a movie by name
  `PUT /api/v1/movies/{name}`

* Delete a movie by name
  `DELETE /api/v1/movies/{name}`

## How to run

- Clone this repository
- Run `go mod download`
- Run `go build`
- Run `./ugly-api`

## TODO

- [x] Logging
- [x] Add Request ID
- [x] User model and User routes
- [ ] JWT Authentication
- [ ] Session
- [ ] Body validation
- [x] Data persistence (MongoDB)
- [x] Handle Errors
- [x] Set environment variables
- [x] Set up a project architecture
- [ ] Unit Tests
- [ ] Docker
- [ ] Swagger