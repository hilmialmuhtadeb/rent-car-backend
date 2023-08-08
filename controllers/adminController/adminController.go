package adminController

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/hilmialmuhtadeb/rent-car-backend/initializers"
	"github.com/hilmialmuhtadeb/rent-car-backend/models"
)

func Create(c *gin.Context) {
	var input struct {
		Email string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(422, gin.H{"error": "Failed to read body request!"})
		return
	}

	var admin models.Admin
	if err := initializers.DB.Where("email = ?", input.Email).First(&admin).Error; err == nil {
		c.JSON(404, gin.H{"error": "Email already used!"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password!"})
		return
	}

	newAdmin := models.Admin{Email: input.Email, Password: string(hash)}
	result := initializers.DB.Create(&newAdmin)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create admin!"})
		return
	}

	c.JSON(200, gin.H{"admin": newAdmin})
}

func Show(c *gin.Context) {
	var admin models.Admin

	token, err := c.Cookie("Authorization")

	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized!"})
		return
	}

	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized!"})
		return
	}

	if err := initializers.DB.Where("id = ?", claims["id"]).First(&admin).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(200, gin.H{"admin": admin})
}

func Delete(c *gin.Context) {
	var admin models.Admin
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&admin).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found!"})
		return
	}

	initializers.DB.Delete(&admin)

	c.JSON(200, gin.H{"message": "Admin deleted!"})
}

func AdminLogin(c *gin.Context) {
	var Body struct {
		Email string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&Body); err != nil {
		c.JSON(422, gin.H{"error": "Failed to read body request!"})
		return
	}

	var admin models.Admin
	if err:= initializers.DB.Where("email = ?", Body.Email).First(&admin).Error; err != nil {
		c.JSON(404, gin.H{"error": "Wrong email or password!"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(Body.Password))

	if err != nil {
		c.JSON(401, gin.H{"error": "Wrong password!"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": admin.Id,
		"email": admin.Email,
		"isAdmin": true,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token!"})
		return
	}

	c.SetCookie("Authorization", tokenString, 3600, "/", "localhost", false, true)
	c.JSON(200, gin.H{"token": tokenString})
}