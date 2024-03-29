package graph

import "github/rowmur/insta-clone/internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DBQueries *database.Queries
}
