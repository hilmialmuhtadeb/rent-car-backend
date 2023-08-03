package carcontroller

import (
	"fmt"
	"path/filepath"

	"github.com/google/uuid"
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
	var input models.CarInput

	if err := c.Bind(&input); err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}
	
	fileExt := filepath.Ext(input.Image.Filename)
	randomFileName := uuid.New().String() + fileExt
	targetPath := fmt.Sprintf("images/%s", randomFileName)

	if err := c.SaveUploadedFile(input.Image, targetPath); err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}

	car := models.Car{Name: input.Name, Type: input.CarType, Rating: input.Rating, Fuel: input.Fuel, Image: randomFileName, HourRate: input.HourRate, DayRate: input.DayRate, MonthRate: input.MonthRate}
	models.DB.Create(&car)

	c.JSON(200, car)
}

func Update(c *gin.Context) {
	var car models.Car
	if err := models.DB.Where("id = ?", c.Param("id")).First(&car).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found!"})
		return
	}

	var input models.CarInput
	if err := c.Bind(&input); err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}

	fileExt := filepath.Ext(input.Image.Filename)
	randomFileName := uuid.New().String() + fileExt
	targetPath := fmt.Sprintf("images/%s", randomFileName)

	if err := c.SaveUploadedFile(input.Image, targetPath); err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}

	updatedCar := models.Car{Name: input.Name, Type: input.CarType, Rating: input.Rating, Fuel: input.Fuel, Image: randomFileName, HourRate: input.HourRate, DayRate: input.DayRate, MonthRate: input.MonthRate}

	models.DB.Model(&car).Updates(updatedCar)

	c.JSON(200, input)
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