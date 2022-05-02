package handlers

import (
	"encoding/json"
	"invite-token/src/models"
	"invite-token/src/services/token"
	"net/http"

	"invite-token/config/constants"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func inviteTokenRoutes(v1 *gin.RouterGroup) {
	v1.POST("/invite-tokens", GenerateInviteTokens)
	v1.PUT("/invalidate/invite-tokens", InValidateInviteToken)
	v1.GET("/invite-tokens", GetAllInviteTokens)

	v1.POST("/validate/invite-tokens", ValidateInviteToken)
}

func InValidateInviteToken(c *gin.Context) {
	if !authenticate(c) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	t := token.New()
	decoder := json.NewDecoder(c.Request.Body)
	req := &models.TokenReq{}

	err := decoder.Decode(req)
	if err != nil {
		c.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	err = t.InValidate(req.Token)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid token",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "token invalidated",
	})
}

func ValidateInviteToken(c *gin.Context) {
	t := token.New()
	decoder := json.NewDecoder(c.Request.Body)
	req := &models.TokenReq{}

	err := decoder.Decode(req)
	if err != nil {
		c.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	valid, err := t.Validate(req.Token)
	if err == constants.ErrInvalidInviteToken {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid token",
		})
		return
	}

	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid token",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"is_valid": valid,
	})
}

func GetAllInviteTokens(c *gin.Context) {
	if !authenticate(c) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	t := token.New()

	tokens, err := t.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tokens": tokens,
	})
}

func GenerateInviteTokens(c *gin.Context) {
	if !authenticate(c) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	t := token.New()

	token, err := t.Generate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"invite_token": token,
	})
}
