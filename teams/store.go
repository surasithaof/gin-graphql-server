package teams

import (
	"context"
	"surasithit/gin-graphql-server/teams/model"
)

type Store interface {
	FindAll(ctx context.Context) ([]*model.Team, error)
	FindOne(ctx context.Context, id int) (*model.Team, error)
	Create(ctx context.Context, team *model.Team) (*model.Team, error)
	Update(ctx context.Context, id int, team *model.Team) (*model.Team, error)
	Delete(ctx context.Context, id int) error
}
