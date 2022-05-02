package handlers

import (
	"invite-token/utils/drivers/db"

	"github.com/gin-gonic/gin"
)

func pingRoutes(router *gin.Engine) {
	router.GET("/ping", Ping)
}

func Ping(c *gin.Context) {
	var no int64
	res := db.DB.Raw(`SELECT 1`).Scan(&no)

	if res.Error != nil {
		c.JSON(500, gin.H{
			"message": res.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
