package graph

import (
	"surasithaof/gin-graphql-server/players"
	"surasithaof/gin-graphql-server/teams"
)

//go:generate go get github.com/99designs/gqlgen
//go:generate go run github.com/99designs/gqlgen generate

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
