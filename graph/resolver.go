package graph

import "graphqltest/database"

//go:generate go run github.com/99designs/gqlgen --verbose

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserRepo     database.UserRepo
	DocumentRepo database.DocumentRepo
}
