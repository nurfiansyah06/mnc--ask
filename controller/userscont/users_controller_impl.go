package userscont

import (
	"golang-test/controller"
	"golang-test/model/users"
	"golang-test/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersContImpl struct {
	userUsecase usecase.UsersUsecase
}

func NewUsersCont(userUsecase usecase.UsersUsecase) controller.UsersController {
	return &UsersContImpl{userUsecase: userUsecase}
}

func (u *UsersContImpl) Register(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := u.userUsecase.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := users.ResponseRegisterSuccess{
		Status: "SUCCESS",
		Result: users.ResponseUser{
			UserID:      createdUser.UserID,
			FirstName:   createdUser.FirstName,
			LastName:    createdUser.LastName,
			PhoneNumber: createdUser.PhoneNumber,
			Address:     createdUser.Address,
			CreatedDate: createdUser.CreatedAt,
		},
	}

	c.JSON(http.StatusOK, response)
}

// Login implements controller.UsersController.
func (u *UsersContImpl) Login(c *gin.Context) {
	var (
		request users.User
	)

	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := u.userUsecase.Login(request.PhoneNumber, request.Pin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := users.ResponseLoginSuccess{
		Status: "SUCCESS",
		Result: users.ResultToken{
			AccessToken:  tokens.AccessToken,
			RefreshToken: tokens.RefreshToken,
		},
	}

	c.JSON(http.StatusOK, response)
}
