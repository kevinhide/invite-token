package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetRouter creates a router and registers all the routes for the
// service and returns it.
func GetRouter() http.Handler {
	router := gin.Default()

	pingRoutes(router)

	v1 := router.Group("/v1")
	v1.Use(CORSMiddleware())
	inviteTokenRoutes(v1)

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
