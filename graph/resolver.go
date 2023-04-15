package graph

import "surasithit/gin-graphql-server/adapters/db"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Database *db.Database
}

func Initialize(database *db.Database) *Resolver {
	return &Resolver{
		Database: database,
	}
}
