package model

import (
	teamModel "surasithit/gin-graphql-server/teams/model"

	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	ID     string          `gorm:"primary_key;auto_increment;not null"`
	Name   string          `gorm:"column:name;not null"`
	Rating float64         `gorm:"column:rating;not null"`
	TeamID int             `gorm:"column:team_id;not null"`
	Team   *teamModel.Team `gorm:"column:team;references:TeamID"`
}