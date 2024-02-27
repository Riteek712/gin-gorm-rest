package controller

import (
	"github.com/Riteek712/gin-gorm-rest/config"
	"github.com/Riteek712/gin-gorm-rest/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	// c.String(200, "hello User Controller.")
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)

}

func CreateUsers(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(200, &user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(200, &user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(200, &user)
}
func GetUserById(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).First(&user)
	c.JSON(200, &user)
}
