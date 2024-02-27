package main

import (
	"github.com/Riteek712/gin-gorm-rest/config"
	"github.com/Riteek712/gin-gorm-rest/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	router.Run(":8000")
}
