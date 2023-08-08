package userController

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/hilmialmuhtadeb/rent-car-backend/initializers"
	"github.com/hilmialmuhtadeb/rent-car-backend/models"
)

func Index(c *gin.Context) {
	var users []models.User
	initializers.DB.Find(&users)

	c.JSON(200, gin.H{"users": users})
}

func Show(c *gin.Context) {
	var user models.User

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

	if err := initializers.DB.Where("id = ?", claims["id"]).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(200, gin.H{"user": user})
}

func Create(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Email string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
		PhoneNumber string `json:"phoneNumber" binding:"required"`
		City string `json:"city" binding:"required"`
		Zip string `json:"zip" binding:"required"`
		Address string `json:"address" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(422, gin.H{"error": "Failed to read body request!"})
		return
	}

	// bad practice, should use unique index instead (somehow it doesn't work)
	var user models.User
	// use ok instead of err
	if err := initializers.DB.Where("email = ?", input.Email).First(&user).Error; err == nil {
		c.JSON(404, gin.H{"error": "Email already used!"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password!"})
		return
	}

	newUser := models.User{Username: input.Username, Email: input.Email, Password: string(hash), PhoneNumber: input.PhoneNumber, City: input.City, Zip: input.Zip, Address: input.Address, Role: 1}
	result := initializers.DB.Create(&newUser)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create user!"})
		return
	}

	c.JSON(200, gin.H{"user": newUser})
}

func Login(c *gin.Context) {
	var Body struct {
		Email string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&Body); err != nil {
		c.JSON(422, gin.H{"error": "Failed to read body request!"})
		return
	}

	var user models.User

	if err := initializers.DB.Where("email = ?", Body.Email).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found!"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Body.Password))

	if err != nil {
		c.JSON(401, gin.H{"error": "Wrong password!"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.Id,
		"email": user.Email,
		"role": user.Role,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token!"})
		return
	}

	c.SetCookie("Authorization", tokenString, 3600, "/", "localhost", false, true)
	c.JSON(200, gin.H{"token": tokenString})
}