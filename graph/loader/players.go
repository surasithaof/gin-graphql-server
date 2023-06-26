package loader

import (
	"context"
	"fmt"
	"log"

	"github.com/surasithaof/gin-graphql-server/players"
	"github.com/surasithaof/gin-graphql-server/players/model"

	"github.com/graph-gophers/dataloader"
)

type PlayerLoader struct {
	store players.Store
}

// Method that implements BatchGetPlayers dataloader.BatchFunc //
func (l *PlayerLoader) BatchGetPlayers(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	// type of dataloader.Keys to []string
	playerIDs := make([]string, len(keys))
	for ix, key := range keys {
		playerIDs[ix] = key.String()
	}
	// Actual process
	playerByID, err := l.store.FindByIDs(ctx, playerIDs)
	if err != nil {
		err := fmt.Errorf("fail get players, %w", err)
		log.Printf("%v \n ", err)
		return nil
	}
	// Convert to []*model.Player[]*dataloader.Result
	output := make([]*dataloader.Result, len(keys))
	for index, key := range keys {
		player, ok := playerByID[key.String()]
		if ok {
			output[index] = &dataloader.Result{Data: player, Error: nil}
		} else {
			err := fmt.Errorf("player not found %s", key.String())
			output[index] = &dataloader.Result{Data: nil, Error: err}
		}
	}
	return output
}

// LoadPlayer dataloader.Load wrapped and typed implementation
func LoadPlayer(ctx context.Context, playerID string) (*model.Player, error) {
	loaders := GetLoaders(ctx)
	thunk := loaders.PlayerLoader.Load(ctx, dataloader.StringKey(playerID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	player := result.(*model.Player)
	return player, nil
}
