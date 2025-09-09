package database

import (
	"fmt"

	"github.com/jakkaphatminthana/go-refresh/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	connectionString := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		config.AppConfig.DatabaseHost,
		config.AppConfig.DatabasePort,
		config.AppConfig.DatabaseUsername,
		config.AppConfig.DatabasePassword,
		config.AppConfig.DatabaseName,
		config.AppConfig.DatabaseSSLMode,
	)

	var initDBErr error
	DB, initDBErr = gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	return DB, initDBErr
}
