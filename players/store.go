package players

import (
	"context"
	"strconv"

	"github.com/surasithaof/gin-graphql-server/players/model"

	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

func Initialize(db *gorm.DB) Store {
	return &Service{
		DB: db,
	}
}

// Implement interface
var _ Store = (*Service)(nil)

// Create implements Store
func (s *Service) Create(ctx context.Context, player *model.Player) (*model.Player, error) {
	tx := s.DB.WithContext(ctx).Create(&player)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return player, nil
}

// Delete implements Store
func (s *Service) Delete(ctx context.Context, id int) error {
	player := &model.Player{}
	tx := s.DB.WithContext(ctx).Delete(&player, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// FindAll implements Store
func (s *Service) FindAll(ctx context.Context) ([]*model.Player, error) {
	players := []*model.Player{}
	tx := s.DB.WithContext(ctx).Find(&players)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return players, nil
}

// FindByIDs implements Store.
func (s *Service) FindByIDs(ctx context.Context, IDs []string) (map[string]*model.Player, error) {

	// func (s *Service) FindByIDs(ctx context.Context, IDs []string) ([]*model.Player, error) {
	players := []*model.Player{}
	tx := s.DB.WithContext(ctx).Where("id IN ?", IDs).Find(&players)
	if tx.Error != nil {
		return nil, tx.Error
	}
	res := make(map[string]*model.Player)
	for _, player := range players {
		id := strconv.Itoa(player.ID)
		res[id] = player
	}

	return res, nil
}

// FindByTeamId implements Store
func (s *Service) FindByTeamId(ctx context.Context, teamId int) ([]*model.Player, error) {
	players := []*model.Player{}
	tx := s.DB.WithContext(ctx).Where("team_id = ?", teamId).Find(&players)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return players, nil
}

// FindByTeamIds implements Store.
func (s *Service) FindByTeamIds(ctx context.Context, teamIds []int) ([]*model.Player, error) {
	players := []*model.Player{}
	tx := s.DB.WithContext(ctx).Where("team_id IN ?", teamIds).Find(&players)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return players, nil
}

// FindOne implements Store
func (s *Service) FindOne(ctx context.Context, id int) (*model.Player, error) {
	player := &model.Player{}
	tx := s.DB.WithContext(ctx).Where("id = ?", id).First(&player)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return player, nil
}

// Update implements Store
func (s *Service) Update(ctx context.Context, id int, player *model.Player) (*model.Player, error) {
	existPlayer, err := s.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}
	updatePlayer := &model.Player{
		ID:     existPlayer.ID,
		Name:   player.Name,
		Rating: player.Rating,
		TeamID: player.TeamID,
	}
	tx := s.DB.WithContext(ctx).Save(&updatePlayer)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return updatePlayer, nil
}
