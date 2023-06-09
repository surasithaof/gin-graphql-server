package teams

import (
	"context"

	"github.com/surasithaof/gin-graphql-server/teams/model"
)

//go:generate mockgen -source=./spec.go -destination=./mocks/store.go -package=mocks "github.com/surasithaof/gin-graphql-server/teams" Store

type Store interface {
	FindAll(ctx context.Context) ([]*model.Team, error)
	FindByIDs(ctx context.Context, IDs []string) (map[string]*model.Team, error)
	FindOne(ctx context.Context, id int) (*model.Team, error)
	Create(ctx context.Context, team *model.Team) (*model.Team, error)
	Update(ctx context.Context, id int, team *model.Team) (*model.Team, error)
	Delete(ctx context.Context, id int) error
}
