package model

type Team struct {
	ID      int    `gorm:"primary_key;auto_increment;not null"`
	Name    string `gorm:"column:name"`
	Country string `gorm:"column:country"`
}
