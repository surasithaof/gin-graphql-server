package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Connection string `envconfig:"DB_CONN"`
}

func Connect(dbConfig Config) *gorm.DB {
	if dbConfig.Connection == "" {
		panic("Not have database connection!")
	}
	db, err := gorm.Open(postgres.Open(dbConfig.Connection), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
