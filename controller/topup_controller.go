package controller

import "github.com/gin-gonic/gin"

type TopupController interface {
	TopUp(c *gin.Context)
}
