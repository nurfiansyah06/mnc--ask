package repository

import "golang-test/model/users"

type UserRepository interface {
	Register(user users.User) (users.User, error)
	Login(phoneNumber, pin string) (*users.User, error)
	SelectByUserId(userId string) (*users.User, error)
}
