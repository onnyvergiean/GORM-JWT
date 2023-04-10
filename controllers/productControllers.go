package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"tugas10/database"
	"tugas10/helpers"
	"tugas10/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

func CreateProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Product := models.Product{}
	userID := uint(userData["id"].(float64))
	User := models.User{}

	err := db.First(&User, "id = ?", userID).Error

	if err != nil {
		c.JSON(http.StatusBadRequest,  gin.H{"error": "Not Found", "message": err.Error()})
		return
	}

	if contentType == appJSON {
		c.BindJSON(&Product)
	} else {
		c.Bind(&Product)
	}

	Product.UserID = userID
	Product.User = &User

	err = db.Create(&Product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, Product)
}

func GetProduct(c *gin.Context) {
	db := database.GetDB()
	Product := models.Product{}
	productId, _ := strconv.Atoi(c.Param("productId"))

	err := db.Preload("User").First(&Product, "id = ?", productId).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Not Found", "message": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
	}
	c.JSON(http.StatusOK, Product)
}


func GetProducts(c *gin.Context) {
	db := database.GetDB()
	Products := []models.Product{}
	err := db.Preload("User").Order("id").Find(&Products).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, Products)
}

func UpdateProduct(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)

	Product := models.Product{}
	productId, _ := strconv.Atoi(c.Param("productId"))
	

	if contentType == appJSON {
		c.BindJSON(&Product)
	} else {
		c.Bind(&Product)
	}
	Product.ID = uint(productId)

	
	err := db.Model(&Product).Where("id = ?", productId).Updates(models.Product{
		Title: Product.Title,
		Description: Product.Description,
	}).Error

	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated", "title": Product.Title, "description": Product.Description})

}


func DeleteProduct(c *gin.Context) {
	db := database.GetDB()
	Product := models.Product{}
	productId, _ := strconv.Atoi(c.Param("productId"))

	err := db.Delete(&Product, "id = ?", productId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}