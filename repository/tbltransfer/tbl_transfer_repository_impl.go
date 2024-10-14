package tbltransfer

import (
	"database/sql"
	"golang-test/model/transfer"
	"golang-test/repository"
	"time"

	"github.com/google/uuid"
)

type transferRepositoryImpl struct {
	db *sql.DB
}

func NewTransferRepository(db *sql.DB) repository.TransferRepository {
	return &transferRepositoryImpl{db}
}

// Transfer implements repository.TransferRepository.
func (t *transferRepositoryImpl) Transfer(userId string, receiveId string, remarks string, amount int) (transfer.Transfer, error) {
	var (
		balanceSender  int
		balanceReceive int
	)

	err := t.db.QueryRow("SELECT balance FROM users WHERE user_id = ?", userId).Scan(&balanceSender)
	if err != nil {
		if err == sql.ErrNoRows {
			return transfer.Transfer{}, err
		}
		return transfer.Transfer{}, err
	}

	err = t.db.QueryRow("SELECT balance FROM users WHERE user_id = ?", receiveId).Scan(&balanceReceive)
	if err != nil {
		if err == sql.ErrNoRows {
			return transfer.Transfer{}, err
		}
		return transfer.Transfer{}, err
	}

	transferID := uuid.New().String()

	_, err = t.db.Exec("UPDATE users SET balance = balance - ? WHERE user_id = ?", amount, userId)
	if err != nil {
		return transfer.Transfer{}, err
	}

	_, err = t.db.Exec("UPDATE users SET balance = balance + ? WHERE user_id = ?", amount, receiveId)
	if err != nil {
		return transfer.Transfer{}, err
	}

	return transfer.Transfer{
		TransferID:    transferID,
		Amount:        amount,
		Remarks:       remarks,
		BalanceBefore: balanceSender,
		BalanceAfter:  balanceReceive,
		CreatedDate:   time.Now(),
	}, nil
}
