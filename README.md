# Calorie Tracker RESTful API

- Gin framework
- Air for hot reloading
- MongoDB for the database

# Routes

## GET

- **/ping/** -> pong [for testing the server]
- **/enteries/** [Get All Enteries]
- **/enteries/:id/** [Get entery with specific id]
- **/ingredients/:ingredient/** [Get entry by the ingredient]

## POST:w

- **/entry/create/** [for creating a new entry]

## PUT

- **entry/update/:id** [Update entry]
- **ingredients/update/:id/** [Update ingredients with entry of a specifc id]

# Project Structure

```bash
├── cmd
│   └── main.go
├── go.mod
├── go.sum
├── LICENSE
├── pkg
│   ├── controllers
│   ├── db
│   │   └── connect.go
│   ├── middleware
│   ├── models
│   │   └── entry.go
│   └── routes
│       └── enteries.go
└── README.md
```
