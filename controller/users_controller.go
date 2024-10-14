package controller

import "github.com/gin-gonic/gin"

type UsersController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}
