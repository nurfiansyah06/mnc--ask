package repository

import "golang-test/model/topup"

type TopUpRepository interface {
	TopUp(topup topup.TopUpRequest, userID string) (topup.TopUpResponseDB, error)
}
