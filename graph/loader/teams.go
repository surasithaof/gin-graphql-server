package loader

import (
	"context"
	"fmt"
	"log"
	"strings"
	"surasithit/gin-graphql-server/teams"
	"surasithit/gin-graphql-server/teams/model"

	"github.com/graph-gophers/dataloader"
)

type TeamLoader struct {
	store teams.Store
}

// Method that implements BatchGetTeams dataloader.BatchFunc //
func (l *TeamLoader) BatchGetTeams(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	// type of dataloader.Keys to []string
	teamIDs := make([]string, len(keys))
	for ix, key := range keys {
		teamIDs[ix] = key.String()
	}
	// Actual process
	log.Printf("BatchGetTeams(id = %s) \n ", strings.Join(teamIDs, ","))
	teamByID, err := l.store.FindByIDs(ctx, teamIDs)
	if err != nil {
		err := fmt.Errorf("fail get teams, %w", err)
		log.Printf("%v \n ", err)
		return nil
	}
	// Convert to []*model.Team[]*dataloader.Result
	output := make([]*dataloader.Result, len(keys))
	for index, key := range keys {
		team, ok := teamByID[key.String()]
		if ok {
			output[index] = &dataloader.Result{Data: team, Error: nil}
		} else {
			err := fmt.Errorf("team not found %s", key.String())
			output[index] = &dataloader.Result{Data: nil, Error: err}
		}
	}
	return output
}

// LoadTeam dataloader.Load wrapped and typed implementation
func LoadTeam(ctx context.Context, teamID string) (*model.Team, error) {
	loaders := GetLoaders(ctx)
	thunk := loaders.TeamLoader.Load(ctx, dataloader.StringKey(teamID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	team := result.(*model.Team)
	return team, nil
}
