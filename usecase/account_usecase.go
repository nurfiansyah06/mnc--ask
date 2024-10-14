package usecase

import "golang-test/model/topup"

type AccountUsecase interface {
	TopUp(topUpRequest topup.TopUpRequest, userId string) (topup.TopUpResponse, error)
}
