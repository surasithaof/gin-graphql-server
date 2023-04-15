package model

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	ID      string `gorm:"primary_key;auto_increment;not null"`
	Name    string `gorm:"column:name"`
	Country string `gorm:"column:country"`
}
