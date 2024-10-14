package tbltopups

import (
	"database/sql"
	"fmt"
	"golang-test/model/topup"
	"golang-test/repository"
	"time"

	"github.com/google/uuid"
)

type topupRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(DB *sql.DB) repository.TopUpRepository {
	return &topupRepositoryImpl{DB: DB}
}

// TopUp implements repository.TopUpRepository.
func (t *topupRepositoryImpl) TopUp(topupRequest topup.TopUpRequest, userID string) (topup.TopUpResponseDB, error) {
	var currentBalance int

	err := t.DB.QueryRow("SELECT balance FROM users WHERE user_id = ?", userID).Scan(&currentBalance)
	if err != nil {
		if err == sql.ErrNoRows {
			return topup.TopUpResponseDB{}, fmt.Errorf("user not found with user_id: %s", userID)
		}
		return topup.TopUpResponseDB{}, fmt.Errorf("failed to get current balance: %w", err)
	}

	topUpID := uuid.New().String()
	newBalance := currentBalance + topupRequest.Amount

	createdDate := time.Now()

	_, err = t.DB.Exec(
		"INSERT INTO top_ups (top_up_id, user_id, amount, balance_before, balance_after, created_date) VALUES (?, ?, ?, ?, ?, ?)",
		topUpID, userID, topupRequest.Amount, currentBalance, newBalance, createdDate,
	)
	if err != nil {
		return topup.TopUpResponseDB{}, fmt.Errorf("failed to insert top-up record: %w", err)
	}

	_, err = t.DB.Exec("UPDATE users SET balance = ? WHERE user_id = ?", newBalance, userID)
	if err != nil {
		return topup.TopUpResponseDB{}, fmt.Errorf("failed to update user's balance: %w", err)
	}

	return topup.TopUpResponseDB{
		TopUpID:       topUpID,
		AmountTopUp:   topupRequest.Amount,
		BalanceBefore: currentBalance,
		BalanceAfter:  newBalance,
		CreatedDate:   createdDate.Format("2006-01-02 15:04:05"),
	}, nil
}
