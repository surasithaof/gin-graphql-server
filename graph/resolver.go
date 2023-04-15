package graph

import "gorm.io/gorm"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Database *gorm.DB
}

func Initialize(database *gorm.DB) *Resolver {
	return &Resolver{
		Database: database,
	}
}
