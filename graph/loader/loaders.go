package loader

import (
	"context"
	"surasithit/gin-graphql-server/players"
	"surasithit/gin-graphql-server/teams"

	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/dataloader"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

// Loaders struct
type Loaders struct {
	PlayerLoader *dataloader.Loader
	TeamLoader   *dataloader.Loader
}

// NewLoaders Loaders initialization method
func NewLoaders(playerStore players.Store, teamStore teams.Store) *Loaders {
	// define the data loader
	userLoader := &PlayerLoader{
		store: playerStore,
	}
	teamLoader := &TeamLoader{
		store: teamStore,
	}
	loaders := &Loaders{
		PlayerLoader: dataloader.NewBatchedLoader(userLoader.BatchGetPlayers),
		TeamLoader:   dataloader.NewBatchedLoader(teamLoader.BatchGetTeams),
	}
	return loaders
}

// HTTP Middleware that injects Middleware Loaders into the context
func Middleware(loaders *Loaders) gin.HandlerFunc {
	loaders.PlayerLoader.ClearAll()
	return func(ctx *gin.Context) {
		nextCtx := context.WithValue(ctx.Request.Context(), loadersKey, loaders)
		ctx.Request = ctx.Request.WithContext(nextCtx)
		ctx.Next()
	}
}

// Get Loaders from GetLoaders Context
func GetLoaders(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
