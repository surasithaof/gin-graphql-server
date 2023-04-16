package teams

import (
	"context"
	"surasithit/gin-graphql-server/teams/model"

	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

func Initialize(db *gorm.DB) Service {
	return Service{
		DB: db,
	}
}

// Implement interface
var _ Store = (*Service)(nil)

// Create implements Store
func (s *Service) Create(ctx context.Context, team *model.Team) (*model.Team, error) {
	tx := s.DB.WithContext(ctx).Create(&team)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return team, nil
}

// Delete implements Store
func (s *Service) Delete(ctx context.Context, id int) error {
	team := &model.Team{}
	tx := s.DB.WithContext(ctx).Delete(&team, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// FindAll implements Store
func (s *Service) FindAll(ctx context.Context) ([]*model.Team, error) {
	teams := []*model.Team{}
	tx := s.DB.WithContext(ctx).Find(&teams)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return teams, nil
}

// FindOne implements Store
func (s *Service) FindOne(ctx context.Context, id int) (*model.Team, error) {
	team := &model.Team{}
	tx := s.DB.WithContext(ctx).Where("id = ?", id).First(&team)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return team, nil
}

// Update implements Store
func (s *Service) Update(ctx context.Context, id int, team *model.Team) (*model.Team, error) {
	existTeam, err := s.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}
	updateTeam := &model.Team{
		ID:      existTeam.ID,
		Name:    team.Name,
		Country: team.Country,
	}
	tx := s.DB.WithContext(ctx).Save(&updateTeam)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return updateTeam, nil
}
