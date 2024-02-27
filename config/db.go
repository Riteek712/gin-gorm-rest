package config

import (
	"github.com/Riteek712/gin-gorm-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("postgresql://postgres:riteek@localhost:5432/Gorm"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	DB = db

}
