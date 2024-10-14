package accounts

import (
	"golang-test/model/topup"
	"golang-test/repository"
	"golang-test/usecase"

	"github.com/go-playground/validator/v10"
)

type AccountUsecaseImpl struct {
	db        repository.TopUpRepository
	validator *validator.Validate
}

func NewAccountUsecase(db repository.TopUpRepository) usecase.AccountUsecase {
	return &AccountUsecaseImpl{
		db:        db,
		validator: validator.New(),
	}
}

// TopUp implements usecase.AccountUsecase.
func (a *AccountUsecaseImpl) TopUp(topUpRequest topup.TopUpRequest, userId string) (topup.TopUpResponse, error) {

	topupResult, err := a.db.TopUp(topUpRequest, userId)
	if err != nil {
		return topup.TopUpResponse{}, err
	}

	result := topup.TopUpResponse{Status: "SUCCESS", Result: topup.TopUpResult{
		TopUpID:       topupResult.TopUpID,
		AmountTopUp:   topupResult.AmountTopUp,
		BalanceBefore: topupResult.BalanceBefore,
		BalanceAfter:  topupResult.BalanceAfter,
		CreatedDate:   topupResult.CreatedDate,
	}}

	return result, nil
}
