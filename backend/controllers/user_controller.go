package controllers

import (
	"area/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMe(c *gin.Context) (string, int) {
	return middlewares.GetClient(c).Hex(), http.StatusOK
}
