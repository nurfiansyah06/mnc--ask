package repository

import "golang-test/model/transfer"

type TransferRepository interface {
	Transfer(userId, receiveId, remarks string, amount int) (transfer.Transfer, error)
}
