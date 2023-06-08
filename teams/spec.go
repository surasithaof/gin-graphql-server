package teams

import (
	"context"
	"surasithit/gin-graphql-server/teams/model"
)

//go:generate mockgen -source=./store.go -destination=./mocks/store.go -package=mocks "surasithit/gin-graphql-server" Store

type Store interface {
	FindAll(ctx context.Context) ([]*model.Team, error)
	FindByIDs(ctx context.Context, IDs []int) ([]*model.Team, error)
	FindOne(ctx context.Context, id int) (*model.Team, error)
	Create(ctx context.Context, team *model.Team) (*model.Team, error)
	Update(ctx context.Context, id int, team *model.Team) (*model.Team, error)
	Delete(ctx context.Context, id int) error
}
