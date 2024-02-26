package db

import (
	"backend/config"
	"backend/models"
	"fmt"
	"log"

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

	if err = db.AutoMigrate(
		&models.Admin{},
		&models.User{},
		&models.Restaurant{},
		&models.Item{},
		&models.Category{},
		&models.Table{},
		&models.Customer{},
		&models.Order{},
		&models.OrderItem{},
	); err != nil {
		log.Fatal("Could not migrate: ", err)
	}

	return db
}
