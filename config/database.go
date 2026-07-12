package config

import (
	"fmt"
	"gin-basic/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func InitDB() *gorm.DB{

	dsn := "host=localhost user=postgres password=postgres dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	fmt.Println("🚀 Database connection successfully established!")
	runMigrations(DB)
	return DB
}

func runMigrations(DB *gorm.DB) {
	err := DB.AutoMigrate(models.GetModels()...)
	if err != nil {
		panic("Failed to run migrations!")
	}
	fmt.Println("✅ Database migrations completed successfully!")
}