package model

import (
	teamModel "surasithaof/gin-graphql-server/teams/model"
)

type Player struct {
	ID     int             `gorm:"primary_key;auto_increment;not null"`
	Name   string          `gorm:"column:name;not null"`
	Rating float64         `gorm:"column:rating;not null"`
	TeamID int             `gorm:"column:team_id;not null"`
	Team   *teamModel.Team `gorm:"foreignKey:TeamID;references:ID"`
}
