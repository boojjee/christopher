package controllers

import (
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(200, gin.H{"status": "OK!", "d": "2930293"})
}
