package graph

import (
	"surasithit/gin-graphql-server/players"
	"surasithit/gin-graphql-server/teams"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	PlayerService players.Service
	TeamService   teams.Service
}

func Initialize(playerService players.Service, teamService teams.Service) *Resolver {
	return &Resolver{
		PlayerService: playerService,
		TeamService:   teamService,
	}
}
