package carcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/hilmialmuhtadeb/rent-car-backend/models"
)

func Index(c *gin.Context) {
	var cars []models.Car
	models.DB.Find(&cars)

	c.JSON(200, gin.H{"cars": cars})
}

func Show(c *gin.Context) {
	var car models.Car

	if err := models.DB.Where("id = ?", c.Param("id")).First(&car).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(200, gin.H{"car": car})
}

func Create(c *gin.Context) {
	var input models.Car

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}

	car := models.Car{Name: input.Name, Type: input.Type, Rating: input.Rating, Fuel: input.Fuel, Image: input.Image, HourRate: input.HourRate, DayRate: input.DayRate, MonthRate: input.MonthRate}
	models.DB.Create(&car)

	c.JSON(200, car)
}

func Update(c *gin.Context) {
	var car models.Car
	if err := models.DB.Where("id = ?", c.Param("id")).First(&car).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found!"})
		return
	}

	var input models.Car
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&car).Updates(input)

	c.JSON(200, car)
}

func Delete(c *gin.Context) {
	var car models.Car
	if err := models.DB.Where("id = ?", c.Param("id")).First(&car).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&car)

	c.JSON(200, gin.H{"success": "Record deleted!"})
}