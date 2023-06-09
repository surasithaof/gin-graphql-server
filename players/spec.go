package players

import (
	"context"

	"github.com/surasithaof/gin-graphql-server/players/model"
)

//go:generate mockgen -source=./spec.go -destination=./mocks/store.go -package=mocks "github.com/surasithaof/gin-graphql-server/players" Store

type Store interface {
	FindAll(ctx context.Context) ([]*model.Player, error)
	FindByIDs(ctx context.Context, IDs []string) (map[string]*model.Player, error)
	FindByTeamId(ctx context.Context, teamId int) ([]*model.Player, error)
	FindByTeamIds(ctx context.Context, teamIds []int) ([]*model.Player, error)
	FindOne(ctx context.Context, id int) (*model.Player, error)
	Create(ctx context.Context, team *model.Player) (*model.Player, error)
	Update(ctx context.Context, id int, team *model.Player) (*model.Player, error)
	Delete(ctx context.Context, id int) error
}
