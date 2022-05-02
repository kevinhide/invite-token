package handlers

import (
	"invite-token/config"

	"github.com/gin-gonic/gin"
)

func authenticate(c *gin.Context) bool {
	uname, password, ok := c.Request.BasicAuth()
	if !ok || uname != config.BasicAuthUserName || password != config.BasicAuthPassword {
		return false
	}

	return true
}
