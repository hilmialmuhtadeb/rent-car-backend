package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hilmialmuhtadeb/rent-car-backend/models"
	"github.com/hilmialmuhtadeb/rent-car-backend/controllers/carcontroller"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/cars", carcontroller.Index)
	r.GET("/api/cars/:id", carcontroller.Show)
	r.POST("/api/cars", carcontroller.Create)
	r.PUT("/api/cars/:id", carcontroller.Update)
	r.DELETE("/api/cars/:id", carcontroller.Delete)

	r.Run()
}