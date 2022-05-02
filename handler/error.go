package handler

import "github.com/gin-gonic/gin"

func Error(c *gin.Context, err error) bool {
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error(), "data": ""})
		return true
	}
	return false
}
