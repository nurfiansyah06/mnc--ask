package usecase

import (
	"golang-test/model/token"
	"golang-test/model/users"
)

type UsersUsecase interface {
	Register(user users.User) (users.User, error)
	Login(phoneNumber, pin string) (token.TokenDetails, error)
}
