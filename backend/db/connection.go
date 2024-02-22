package db

import (
	"backend/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(cfg *config.Config) *gorm.DB {
	dataSourceName := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DB.Host,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Name,
		cfg.DB.Port,
	)

	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return db
}
