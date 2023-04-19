package graph

import (
	"surasithit/gin-graphql-server/players"
	"surasithit/gin-graphql-server/teams"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	PlayerStore players.Store
	TeamStore   teams.Store
}

func Initialize(playerStore players.Store, teamStore teams.Store) *Resolver {
	return &Resolver{
		PlayerStore: playerStore,
		TeamStore:   teamStore,
	}
}
