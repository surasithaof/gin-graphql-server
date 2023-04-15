package db

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Connection string `envconfig:"DB_CONN"`
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
	// db.Logger = logger.Default.LogMode(logger.Info)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate()
	return db
}
