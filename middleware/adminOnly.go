package middleware

import (
	"fmt"
	"os"
	
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AdminOnly(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized!"})
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
	
		return []byte(os.Getenv("SECRET")), nil
	})
	
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}
	
	
	c.Next()
}