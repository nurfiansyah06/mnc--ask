package controller

import (
	"github.com/gin-gonic/gin"
)

type TransferController interface {
	Transfer(c *gin.Context)
}
