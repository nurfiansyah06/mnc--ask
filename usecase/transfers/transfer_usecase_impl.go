package transfers

import (
	"fmt"
	"golang-test/model/transfer"
	"golang-test/repository"
	"golang-test/usecase"

	"github.com/go-playground/validator/v10"
)

type TransferUsecaseImpl struct {
	db        repository.TransferRepository
	userRepo  repository.UserRepository
	validator *validator.Validate
}

func NewTransferUsecase(db repository.TransferRepository, userRepo repository.UserRepository) usecase.TransferUsecase {
	return &TransferUsecaseImpl{
		db:        db,
		validator: validator.New(),
		userRepo:  userRepo,
	}
}

// Transfer implements usecase.TransferUsecase.
func (t *TransferUsecaseImpl) Transfer(userId string, receiveId string, remarks string, amount int) (transfer.Transfer, error) {

	userBalance, err := t.userRepo.SelectByUserId(userId)
	if err != nil {
		return transfer.Transfer{}, err
	}
	if userBalance.Balance < amount {
		return transfer.Transfer{}, fmt.Errorf("balance not enough")
	}

	transfer, err := t.db.Transfer(userId, receiveId, remarks, amount)
	if err != nil {
		return transfer, err
	}

	return transfer, nil

}
