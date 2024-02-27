package routes

import (
	"github.com/Riteek712/gin-gorm-rest/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/users", controller.GetUsers)
	router.POST("/users", controller.CreateUsers)
	router.GET("/users/:id", controller.GetUserById)
	router.DELETE("/users/:id", controller.DeleteUser)
	router.PUT("/users/:id", controller.UpdateUser)

}
