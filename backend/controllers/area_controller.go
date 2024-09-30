package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAREA(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "AREA created successfully"})
}
