package usecase

import (
	"golang-test/model/transfer"
)

type TransferUsecase interface {
	Transfer(userId string, receiveId string, remarks string, amount int) (transfer.Transfer, error)
}
