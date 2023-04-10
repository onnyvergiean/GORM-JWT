package controllers

import (
	"net/http"
	"tugas10/database"
	"tugas10/helpers"
	"tugas10/models"

	"github.com/gin-gonic/gin"
)

var appJSON = "application/json"

func UserRegister(c *gin.Context){
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	

	if contentType == appJSON {
		c.BindJSON(&User)
	} else {
		c.Bind(&User)
	}

	err := db.Create(&User).Error
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": User.ID, "full_name": User.FullName, "email": User.Email, "role": User.Role})
}

func UserLogin(c *gin.Context){
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.BindJSON(&User)
	} else {
		c.Bind(&User)
	}

	password = User.Password
	err := db.Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	comparePassword := helpers.ComparePassword([]byte(User.Password), []byte(password))

	if !comparePassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Password is not match"})
		return
	}

	token, err := helpers.GenerateToken(User.ID, User.Email, User.Role)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}