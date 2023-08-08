package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/hilmialmuhtadeb/rent-car-backend/controllers/adminController"
	"github.com/hilmialmuhtadeb/rent-car-backend/controllers/carController"
	"github.com/hilmialmuhtadeb/rent-car-backend/controllers/orderController"
	"github.com/hilmialmuhtadeb/rent-car-backend/controllers/userController"

	"github.com/hilmialmuhtadeb/rent-car-backend/middleware"

	"github.com/hilmialmuhtadeb/rent-car-backend/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	r.Static("/images", "./images")

	r.GET("/api/cars", carController.Index)
	r.GET("/api/cars/:id", carController.Show)
	r.POST("/api/cars", carController.Create)
	r.PUT("/api/cars/:id", carController.Update)
	r.DELETE("/api/cars/:id", carController.Delete)

	r.GET("/api/users", userController.Index)
	r.GET("/api/users/profile", userController.Show)
	r.POST("/api/users", userController.Create)
	r.POST("/api/login", userController.Login)

	r.GET("/api/admin", adminController.Show)
	r.POST("/api/admin/login", adminController.AdminLogin)
	r.POST("/api/admin", adminController.Create)
	r.DELETE("/api/admin/:id", adminController.Delete)

	r.GET("/api/orders", middleware.AdminOnly, orderController.Index)
	r.GET("/api/orders/:id", orderController.Show)
	r.POST("/api/orders", orderController.Create)
	r.PUT("/api/orders/:id", orderController.Update)
	r.DELETE("/api/orders/:id", orderController.Delete)

	r.Run()
}