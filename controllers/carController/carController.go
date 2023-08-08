package carController

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hilmialmuhtadeb/rent-car-backend/initializers"
	"github.com/hilmialmuhtadeb/rent-car-backend/models"
)

func Index(c *gin.Context) {
	var cars []models.Car
	initializers.DB.Find(&cars)

	c.JSON(200, gin.H{"cars": cars})
}

func Show(c *gin.Context) {
	var car models.Car

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&car).Error; err != nil {
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
	initializers.DB.Create(&car)

	c.JSON(200, gin.H{"car": car})
}

func Update(c *gin.Context) {
	var car models.Car
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&car).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found!"})
		return
	}

	var input models.CarInput
	if err := c.Bind(&input); err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}

	var imageFilePath string
	oldFilePath := fmt.Sprintf("images/%s", car.Image)

	// Check if the user sent a new file in the update request
	_, fileHeader, err := c.Request.FormFile("image")
	if err == nil && fileHeader != nil {
		// The user sent a new file, perform the update logic and save the new file
		// ...
		fileExt := filepath.Ext(fileHeader.Filename)
		randomFileName := uuid.New().String() + fileExt
		targetPath := fmt.Sprintf("images/%s", randomFileName)

		imageFilePath = randomFileName

		if err := c.SaveUploadedFile(input.Image, targetPath); err != nil {
			c.JSON(422, gin.H{"error": err.Error()})
			return
		}
		// After the update, remove the old file
		err := os.Remove(oldFilePath)
		if err != nil {
			fmt.Println("Failed to remove the old file:", err)
		}
	} else {
		imageFilePath = car.Image
	}


	updatedCar := models.Car{Name: input.Name, Type: input.CarType, Rating: input.Rating, Fuel: input.Fuel, Image: imageFilePath, HourRate: input.HourRate, DayRate: input.DayRate, MonthRate: input.MonthRate}

	initializers.DB.Model(&car).Updates(updatedCar)

	c.JSON(200, updatedCar)
}

func Delete(c *gin.Context) {
	var car models.Car
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&car).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found!"})
		return
	}
	
	initializers.DB.Delete(&car)

	filePath := fmt.Sprintf("images/%s", car.Image)

	err := os.Remove(filePath)
	if err != nil {
		fmt.Println("Failed to remove the old file:", err)
	}


	c.JSON(200, gin.H{"success": "Record deleted!"})
}