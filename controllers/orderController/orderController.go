package orderController

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/hilmialmuhtadeb/rent-car-backend/initializers"
	"github.com/hilmialmuhtadeb/rent-car-backend/models"
)

func Index(c *gin.Context) {
	var orders []models.Order
	initializers.DB.Find(&orders)

	c.JSON(200, gin.H{"orders": orders})
}

func Show(c *gin.Context) {
	var order models.Order

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(200, gin.H{"order": order})
}

func Create(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized!"})
		return
	}

	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})

	var input models.OrderInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}

	userID := int64(claims["id"].(float64))

	var order models.Order = models.Order{CarId: input.CarId, UserId: userID, AdminId: 1, PickupLocation: input.PickupLocation, DropoffLocation: input.DropoffLocation, PickupDate: input.PickupDate, DropoffDate: input.DropoffDate, PickupTime: input.PickupTime}
	
	initializers.DB.Create(&order)

	c.JSON(200, gin.H{"order": order})
}

func Update(c *gin.Context) {
	var order models.Order
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found!"})
		return
	}

	var input models.OrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Model(&order).Updates(input)

	c.JSON(200, gin.H{"order": order})
}

func Delete(c *gin.Context) {
	var order models.Order
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found!"})
		return
	}

	initializers.DB.Delete(&order)

	c.JSON(200, gin.H{"success": "Order deleted successfully"})
}