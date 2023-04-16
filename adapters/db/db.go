package db

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	playerModel "surasithit/gin-graphql-server/players/model"
	teamModel "surasithit/gin-graphql-server/teams/model"
)

type Config struct {
	Connection  string `envconfig:"DB_CONN" required:"true"`
	AutoMigrate bool   `envconfig:"DB_AUTO_MIGRATE" default:"false"`
}

func Connect(dbConfig Config) *gorm.DB {
	if dbConfig.Connection == "" {
		panic("Not have database connection!")
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(postgres.Open(dbConfig.Connection), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	if dbConfig.AutoMigrate {
		db.AutoMigrate(
			&teamModel.Team{},
			&playerModel.Player{},
		)
	}

	return db
}
