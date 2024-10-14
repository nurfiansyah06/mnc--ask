package users

import (
	"golang-test/middleware"
	"golang-test/model/token"
	"golang-test/model/users"
	"golang-test/repository"
	"golang-test/usecase"

	"github.com/go-playground/validator/v10"
)

type UsersUsecaseImpl struct {
	db        repository.UserRepository
	validator *validator.Validate
}

func NewUsersUsecase(db repository.UserRepository) usecase.UsersUsecase {
	return &UsersUsecaseImpl{
		db:        db,
		validator: validator.New(),
	}
}

func (u *UsersUsecaseImpl) Register(user users.User) (users.User, error) {
	err := u.validator.Struct(user)
	if err != nil {
		return user, err
	}

	createdUser, err := u.db.Register(user)
	if err != nil {
		return user, err
	}

	return createdUser, nil
}

// Login implements usecase.UsersUsecase.
func (u *UsersUsecaseImpl) Login(phoneNumber, pin string) (token.TokenDetails, error) {
	user, err := u.db.Login(phoneNumber, pin)
	if err != nil {
		return token.TokenDetails{}, err
	}

	if user.Pin != pin {
		return token.TokenDetails{}, err
	}

	accessToken, err := middleware.GenerateTokens(user.UserID)
	if err != nil {
		return token.TokenDetails{}, err
	}

	refreshToken, err := middleware.GenerateRefreshToken(user.UserID)
	if err != nil {
		return token.TokenDetails{}, err
	}

	resultToken := token.TokenDetails{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return resultToken, nil
}
