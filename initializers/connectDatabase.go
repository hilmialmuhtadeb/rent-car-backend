package initializers

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/hilmialmuhtadeb/rent-car-backend/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root@tcp(localhost:3306)/rent-car-app"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.Car{}, &models.User{}, &models.Order{}, &models.Admin{})

	DB = database
}