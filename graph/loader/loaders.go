package loader

import (
	"context"

	"github.com/surasithaof/gin-graphql-server/players"
	"github.com/surasithaof/gin-graphql-server/teams"

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

	loaderOptions := []dataloader.Option{
		dataloader.WithClearCacheOnBatch(),
		// dataloader.WithWait(10 * time.Millisecond),
		// dataloader.WithBatchCapacity(100),
	}

	loaders := &Loaders{
		PlayerLoader: dataloader.NewBatchedLoader(
			userLoader.BatchGetPlayers,
			loaderOptions...,
		),
		TeamLoader: dataloader.NewBatchedLoader(
			teamLoader.BatchGetTeams,
			loaderOptions...,
		),
	}
	return loaders
}

// HTTP Middleware that injects Middleware Loaders into the context
func Middleware(loaders *Loaders) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		loaders.PlayerLoader.ClearAll()
		loaders.TeamLoader.ClearAll()

		nextCtx := context.WithValue(ctx.Request.Context(), loadersKey, loaders)
		ctx.Request = ctx.Request.WithContext(nextCtx)
		ctx.Next()
	}
}

// Get Loaders from GetLoaders Context
func GetLoaders(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
