package initializers

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"

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