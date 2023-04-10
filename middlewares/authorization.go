package middlewares

import (
	"net/http"
	"strconv"
	"tugas10/database"
	"tugas10/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		productId, err := strconv.Atoi(c.Param("productId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message":"Invalid Parameter"})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Product := models.Product{}
	
	
		err = db.Select("user_id").First(&Product, uint(productId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Not Found", "message": "Product not found"})
			return
		}
		
		if userData["role"] != "admin" && Product.UserID != userID {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden", "message": "You are not authorized to access this data"})
				return	
		}

		c.Next()
	}
}

func UserAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		User := models.User{}

		err := db.First(&User, "id = ?", userID).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Not Found", "message": "User not found"})
			return
		}

		if User.Role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden", "message": "You are not authorized to access this endpoint"})
			return
		}

		c.Next()
	}
}